package v1

import (
	"context"

	analyticV1 "analytic-service/internal/pkg/pb/analytic-service/analytic/v1"
)

func (i *Implementation) GetUserTaskCount(ctx context.Context, req *analyticV1.GetUserTaskCountRequest) (*analyticV1.GetUserTaskCountResponse, error) {
	count, err := i.services.GetTaskCount.CountTasks(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &analyticV1.GetUserTaskCountResponse{
		TaskCount: count,
	}, nil
}
