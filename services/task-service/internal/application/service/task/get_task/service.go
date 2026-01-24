package get_task

import (
	"context"

	"task-service/internal/domain/entity"

	"github.com/gofrs/uuid"
)

type TaskProvider interface {
	GetTaskByID(ctx context.Context, id uuid.UUID) (*entity.Task, error)
}

type Service struct {
	taskProvider TaskProvider
}

func NewService(taskProvider TaskProvider) *Service {
	return &Service{taskProvider: taskProvider}
}

func (s *Service) GetTask(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	return s.taskProvider.GetTaskByID(ctx, id)
}
