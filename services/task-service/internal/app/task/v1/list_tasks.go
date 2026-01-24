package v1

import (
	"context"

	"task-service/internal/domain/entity"
	"task-service/internal/pkg/convert"
	taskV1 "task-service/internal/pkg/pb/task-service/task/v1"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/structpb"
)

func (i *Implementation) ListTasks(ctx context.Context, req *taskV1.ListTasksRequest) (*taskV1.ListTasksResponse, error) {
	tasks, err := i.services.ListTasks.ListTasks(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &taskV1.ListTasksResponse{
		Tasks: lo.Map(tasks, func(task *entity.Task, _ int) *taskV1.ListTasksResponse_Task {
			result, _ := structpb.NewStruct(task.Result)
			return &taskV1.ListTasksResponse_Task{
				TaskId: task.ID.String(),
				Status: string(task.Status),
				Price:  convert.DecimalToMoney(task.Price),
				Result: result,
			}
		}),
	}, nil
}
