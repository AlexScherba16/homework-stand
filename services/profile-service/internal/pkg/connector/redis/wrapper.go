package redis

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync/atomic"

	"profile-service/config"
	"profile-service/internal/pkg/connector/redis/hashtag"
	"profile-service/internal/pkg/connector/redis/slot"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

type ShardedClient struct {
	*redis.ClusterClient
	// диапазоны слотов (shards)
	shards slot.Slice
}

// shardByKey возвращает ID шарда для ключа
func (s *ShardedClient) shardByKey(key string) (int, error) {
	hashed := hashtag.Slot(key)
	for _, r := range s.shards {
		if hashed >= r.Start && hashed <= r.End {
			return r.ShardID, nil
		}
	}
	return 0, fmt.Errorf("no shard for slot %d", hashed)
}

func NewShardedClient(ctx context.Context) (*ShardedClient, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        config.Instance().RedisCluster.Addresses,
		DialTimeout:  config.Instance().RedisCluster.DialTimeout,
		MaxRetries:   config.Instance().RedisCluster.MaxRetries,
		ReadTimeout:  config.Instance().RedisCluster.ReadTimeout,
		WriteTimeout: config.Instance().RedisCluster.WriteTimeout,
	})

	err := client.ForEachShard(ctx, func(ctx context.Context, shClient *redis.Client) error {
		return shClient.Ping(ctx).Err()
	})

	if err != nil {
		return nil, err
	}

	// получаем "топологию"
	slots, err := client.ClusterSlots(ctx).Result()
	if err != nil {
		return nil, err
	}

	var shards = make(slot.Slice, 0, len(slots))

	// формируем маппинг
	for shardIdx, slotRange := range slots {
		shards = append(shards, &slot.Range{
			Start:   slotRange.Start,
			End:     slotRange.End,
			ShardID: shardIdx,
		})
	}

	sort.Sort(shards)

	return &ShardedClient{
		shards:        shards,
		ClusterClient: client,
	}, nil
}

func (s *ShardedClient) PartialMGet(ctx context.Context, keys ...string) *StrSliceCmd {
	type keysPositions struct {
		keys []string
		pos  []int
	}

	cmdResult := &StrSliceCmd{}

	var errorKeys int32
	results := make([]interface{}, len(keys))

	shardToKeysPos := make(map[int]keysPositions, min(len(s.shards), len(keys)))

	for position, key := range keys {
		shardID, err := s.shardByKey(key)
		if err != nil {
			errorKeys++
			results[position] = err
			continue
		}

		tmp, ok := shardToKeysPos[shardID]
		if !ok {
			tmp.keys = make([]string, 0)
			tmp.pos = make([]int, 0)
		}

		tmp.keys = append(tmp.keys, key)
		tmp.pos = append(tmp.pos, position)

		shardToKeysPos[shardID] = tmp
	}

	wg, egCtx := errgroup.WithContext(ctx)

	// не отменяем запрос при первой ошибке для PartialMGet
	egCtx = ctx

	// проходимся по сгруппированным ключам
	for _, keysPos := range shardToKeysPos {
		wg.Go(func() error {
			cmd := s.crossSlot(egCtx, keysPos.keys...)
			if cmd.Err() != nil {

				for idx := range keysPos.keys {
					results[keysPos.pos[idx]] = cmd.Err()
				}

				atomic.AddInt32(&errorKeys, int32(len(keysPos.keys)))
				return cmd.Err()
			}

			for idx := range cmd.Val() {
				results[keysPos.pos[idx]] = cmd.Val()[idx]
			}
			return nil
		})
	}

	err := wg.Wait()
	successKeys := len(keys) - int(errorKeys)
	if err != nil || errorKeys > 0 {
		if successKeys <= 0 {
			cmdResult.err = err
			return cmdResult
		}

		cmdResult.err = &PartialError{
			SuccessKeys: successKeys,
			ErrorKeys:   int(errorKeys),
		}
	}

	cmdResult.val = results
	return cmdResult
}

func (s *ShardedClient) crossSlot(ctx context.Context, keys ...string) *StrSliceCmd {
	pipeline := s.ClusterClient.Pipeline()

	commandList := make([]*redis.StringCmd, len(keys))
	for _, key := range keys {
		commandList = append(commandList, pipeline.Get(ctx, key))
	}

	cmds, err := pipeline.Exec(ctx)
	if err != nil && !errors.Is(err, redis.Nil) {
		return &StrSliceCmd{err: err}
	}

	var values = make([]interface{}, 0, len(cmds))

	for _, cmd := range commandList {
		val, errR := cmd.Result()
		if errR != nil && errors.Is(errR, redis.Nil) {
			values = append(values, nil)
			continue
		}

		if errR != nil {
			return &StrSliceCmd{err: errR}
		}
		values = append(values, val)
	}

	return &StrSliceCmd{val: values}
}
