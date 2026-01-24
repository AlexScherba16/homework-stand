package storage

import (
	"profile-service/internal/infrastructure/storage/profile"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Registry struct {
	Profile *profile.Storage
}

func NewRegistry(pool *pgxpool.Pool) *Registry {
	return &Registry{
		Profile: profile.NewStorage(pool),
	}
}
