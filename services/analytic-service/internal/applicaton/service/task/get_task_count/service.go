package get_task_count

import (
	"context"
)

type Storage interface {
	CountTasks(ctx context.Context, userID int64) (uint64, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) CountTasks(ctx context.Context, userID int64) (uint64, error) {
	return s.storage.CountTasks(ctx, userID)
}
