package v1

import (
	"context"

	taskV1 "api-gateway/internal/pkg/pb/api-gateway/task/v1"
	externalV1 "api-gateway/internal/pkg/pb/external/task-service/task/v1"
)

func (i *Implementation) CreateTask(ctx context.Context, req *taskV1.CreateTaskRequest) (*taskV1.CreateTaskResponse, error) {
	response, err := i.external.CreateTask(ctx, &externalV1.CreateTaskRequest{
		UserId:        req.GetUserId(),
		CategoryId:    req.GetCategoryId(),
		Comment:       req.GetComment(),
		ExecutionTime: req.GetExecutionTime(),
	})
	if err != nil {
		return nil, err
	}

	return &taskV1.CreateTaskResponse{
		TaskId: response.GetTaskId(),
		Status: response.GetStatus(),
		Price:  response.GetPrice(),
	}, nil
}
