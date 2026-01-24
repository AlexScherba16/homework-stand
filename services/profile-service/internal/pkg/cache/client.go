package cache

import (
	"context"

	pkgredis "profile-service/internal/pkg/connector/redis"

	"github.com/pkg/errors"
)

// start lifetime version
const startFrom = 1

type Client[TValue Marshaller, TValuePtr UnMarshaller] struct {
	client *pkgredis.ShardedClient
}

func NewClient[TValue Marshaller, TValuePtr UnMarshaller](client *pkgredis.ShardedClient) *Client[TValue, TValuePtr] {
	return &Client[TValue, TValuePtr]{
		client: client,
	}
}

//  ------------------  GET/SET ------------------

func (c *Client[TValue, TValuePtr]) Set(ctx context.Context, entry Entry[TValue, TValuePtr]) error {
	return c.client.Set(ctx, entry.Key, entry.marshall(), entry.TTL()).Err()
}

func (c *Client[TValue, TValuePtr]) MGet(ctx context.Context, keys ...string) (Entries[TValue, TValuePtr], []string, error) {
	var partErr *pkgredis.PartialError
	// получаем частичный результат из кеша
	cmd := c.client.PartialMGet(ctx, keys...)

	if cmd.Err() != nil && !errors.As(cmd.Err(), &partErr) {
		return nil, keys, cmd.Err()
	}

	values := cmd.Val()
	entries := make(Entries[TValue, TValuePtr], 0, len(values))
	missed := make([]string, 0)

	for i, raw := range values {
		key := keys[i]
		switch asserted := raw.(type) {
		case error, nil:
			missed = append(missed, key)
			continue
		case string:
			if len(asserted) == 0 {
				continue
			}

			entry, err := From[TValue, TValuePtr](key, []byte(asserted))
			if err != nil {
				missed = append(missed, key)
				continue
			}

			entries = append(entries, entry)
		}
	}

	return entries, missed, nil
}
