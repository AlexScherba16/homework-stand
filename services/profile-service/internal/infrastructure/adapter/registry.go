package adapter

import (
	"profile-service/internal/infrastructure/adapter/profile"
	"profile-service/internal/infrastructure/adapter/profile/payload"
	"profile-service/internal/infrastructure/storage"
	"profile-service/internal/pkg/cache"
	pkgredis "profile-service/internal/pkg/connector/redis"
)

type Registry struct {
	Profile *profile.Adapter
}

func NewRegistry(redis *pkgredis.ShardedClient, dal *storage.Registry) *Registry {
	return &Registry{
		Profile: profile.NewAdapter(cache.NewClient[payload.Profile, *payload.Profile](redis), dal.Profile),
	}
}
