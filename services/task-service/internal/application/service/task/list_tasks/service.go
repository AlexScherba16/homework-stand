package list_tasks

import (
	"context"

	"task-service/internal/domain/entity"
)

type TaskProvider interface {
	GetTasksByUserID(ctx context.Context, userID int64) ([]*entity.Task, error)
}

type Service struct {
	taskProvider TaskProvider
}

func NewService(taskProvider TaskProvider) *Service {
	return &Service{taskProvider: taskProvider}
}

func (s *Service) ListTasks(ctx context.Context, userID int64) ([]*entity.Task, error) {
	return s.taskProvider.GetTasksByUserID(ctx, userID)
}
