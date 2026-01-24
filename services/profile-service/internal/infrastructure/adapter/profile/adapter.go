package profile

import (
	"context"
	"fmt"
	"log/slog"

	"profile-service/config"
	"profile-service/internal/domain/entity"
	"profile-service/internal/infrastructure/adapter/profile/payload"
	"profile-service/internal/pkg/cache"

	"github.com/samber/lo"
)

type Entry = cache.Entry[payload.Profile, *payload.Profile]

type Storage interface {
	CreateProfile(ctx context.Context, p *entity.Profile) error
	GetProfiles(ctx context.Context, userIDs []int64) ([]*entity.Profile, error)
}

type Adapter struct {
	client  *cache.Client[payload.Profile, *payload.Profile]
	storage Storage
}

func NewAdapter(client *cache.Client[payload.Profile, *payload.Profile], origin Storage) *Adapter {
	return &Adapter{
		client:  client,
		storage: origin,
	}
}

func (a *Adapter) CreateProfile(ctx context.Context, profile *entity.Profile) error {
	err := a.storage.CreateProfile(ctx, profile)
	if err != nil {
		return err
	}

	// кладем в кеш, но нам тут не хочется делать CAS и прочее :) (ради простоты, ибо задача в другом)
	return a.client.Set(ctx, Entry{
		Key:        payload.ProfileKey(profile.UserID, config.Instance().Cache.KeyVersion),
		Value:      payload.ConvertFrom(profile),
		Expiration: config.Instance().Cache.TTL,
	})
}

func (a *Adapter) GetProfiles(ctx context.Context, userIDs ...int64) (map[int64]*entity.Profile, error) {
	// формируем ключи
	keys := lo.Map(userIDs, func(userID int64, _ int) string {
		return payload.ProfileKey(userID, config.Instance().Cache.KeyVersion)
	})

	found, missed, err := a.client.MGet(ctx, keys...)
	if err != nil {
		// если есть промах в кеш по причине ошибки
		slog.Error(fmt.Sprintf("error from client.MGet: %s", err.Error()))
		missed = keys
	}

	var result = make(map[int64]*entity.Profile, len(userIDs))
	for _, entry := range found {
		slog.Info(fmt.Sprintf("found key: %s", entry.Key))
		result[entry.Value.UserID] = entry.Value.ConvertTo()
	}

	if len(missed) == 0 {
		return result, nil
	}

	slog.Info(fmt.Sprintf("read missed '%v' from database", missed))

	profiles, err := a.storage.GetProfiles(ctx, lo.Map(missed, func(key string, _ int) int64 {
		return payload.ProfileID(key)
	}))
	if err != nil {
		return nil, err
	}

	for _, profile := range profiles {
		result[profile.UserID] = profile
	}

	return result, nil
}
