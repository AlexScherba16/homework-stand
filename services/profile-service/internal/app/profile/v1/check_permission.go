package v1

import (
	"context"

	profileV1 "profile-service/internal/pkg/pb/profile-service/profile/v1"
)

func (i *Implementation) CheckPermission(ctx context.Context, req *profileV1.CheckPermissionRequest) (*profileV1.CheckPermissionResponse, error) {
	allowed, err := i.services.CheckPermission.CheckPermission(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &profileV1.CheckPermissionResponse{
		Allowed: allowed,
	}, nil
}
