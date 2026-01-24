package storage

import (
	"task-service/internal/infrastructure/storage/category"
	"task-service/internal/infrastructure/storage/task"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Registry struct {
	Category *category.Storage
	Task     *task.Storage
}

func NewRegistry(pool *pgxpool.Pool) *Registry {
	return &Registry{
		Category: category.NewStorage(),
		Task:     task.NewStorage(pool),
	}
}
