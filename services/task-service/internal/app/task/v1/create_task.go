package v1

import (
	"context"

	"task-service/internal/application/service/task/create_task"
	"task-service/internal/pkg/convert"
	taskV1 "task-service/internal/pkg/pb/task-service/task/v1"
)

func (i *Implementation) CreateTask(ctx context.Context, req *taskV1.CreateTaskRequest) (*taskV1.CreateTaskResponse, error) {
	dto := create_task.CreateTaskDTO{
		UserID:        req.GetUserId(),
		CategoryID:    req.GetCategoryId(),
		Comment:       req.GetComment(),
		ExecutionTime: req.GetExecutionTime().AsDuration(),
	}

	task, err := i.services.CreateTask.Create(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &taskV1.CreateTaskResponse{
		TaskId: task.ID.String(),
		Status: string(task.Status),
		Price:  convert.DecimalToMoney(task.Price),
	}, nil
}
