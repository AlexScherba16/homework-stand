package v1

import (
	"context"

	taskV1 "api-gateway/internal/pkg/pb/api-gateway/task/v1"
	externalV1 "api-gateway/internal/pkg/pb/external/task-service/task/v1"

	"github.com/samber/lo"
)

func (i *Implementation) ListTasks(ctx context.Context, req *taskV1.ListTasksRequest) (*taskV1.ListTasksResponse, error) {
	response, err := i.external.ListTasks(ctx, &externalV1.ListTasksRequest{
		UserId: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}

	return &taskV1.ListTasksResponse{
		Tasks: lo.Map(response.GetTasks(), func(task *externalV1.ListTasksResponse_Task, _ int) *taskV1.ListTasksResponse_Task {
			return &taskV1.ListTasksResponse_Task{
				TaskId: task.GetTaskId(),
				Status: task.GetStatus(),
				Price:  task.GetPrice(),
				Result: task.GetResult(),
			}
		}),
	}, nil
}
