package storage

import (
	"analytic-service/internal/infrastructure/storage/task"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Registry struct {
	TaskStorage *task.Storage
}

func NewRegistry(pool *pgxpool.Pool) *Registry {
	return &Registry{
		TaskStorage: task.NewStorage(pool),
	}
}
