package v1

import (
	"context"

	taskV1 "api-gateway/internal/pkg/pb/api-gateway/task/v1"
	externalV1 "api-gateway/internal/pkg/pb/external/task-service/task/v1"
)

func (i *Implementation) GetTask(ctx context.Context, req *taskV1.GetTaskRequest) (*taskV1.GetTaskResponse, error) {
	response, err := i.external.GetTask(ctx, &externalV1.GetTaskRequest{
		TaskId: req.GetTaskId(),
	})
	if err != nil {
		return nil, err
	}

	return &taskV1.GetTaskResponse{
		TaskId: response.GetTaskId(),
		Status: response.GetStatus(),
		Price:  response.GetPrice(),
		Result: response.GetResult(),
	}, nil
}
