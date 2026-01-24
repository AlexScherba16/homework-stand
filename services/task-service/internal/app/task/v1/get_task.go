package v1

import (
	"context"

	"task-service/internal/pkg/convert"
	taskV1 "task-service/internal/pkg/pb/task-service/task/v1"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/structpb"
)

func (i *Implementation) GetTask(ctx context.Context, req *taskV1.GetTaskRequest) (*taskV1.GetTaskResponse, error) {
	task, err := i.services.GetTask.GetTask(ctx, uuid.FromStringOrNil(req.GetTaskId()))
	if err != nil {
		return nil, err
	}

	result, err := structpb.NewStruct(task.Result)
	if err != nil {
		return nil, err
	}

	return &taskV1.GetTaskResponse{
		TaskId: task.ID.String(),
		Status: string(task.Status),
		Price:  convert.DecimalToMoney(task.Price),
		Result: result,
	}, nil
}
