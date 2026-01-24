package accept_task

import (
	"context"

	"analytic-service/internal/domain/entity"
)

type Creator interface {
	InsertTask(ctx context.Context, task *entity.Task) error
}

type Service struct {
	storage Creator
}

func NewService(storage Creator) *Service {
	return &Service{storage: storage}
}

func (s *Service) Create(ctx context.Context, request CreateTaskRequest) error {
	task := entity.NewTask(
		request.taskID,
		request.userID,
		request.categoryID,
		request.status,
		request.comment,
		request.executionTime,
		request.createdAt,
		request.price,
	)

	return s.storage.InsertTask(ctx, task)
}
