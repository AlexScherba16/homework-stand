package v1

import (
	"context"

	"profile-service/internal/application/service/profile/create_profile"
	profileV1 "profile-service/internal/pkg/pb/profile-service/profile/v1"
)

func (i *Implementation) CreateProfile(ctx context.Context, req *profileV1.CreateProfileRequest) (*profileV1.CreateProfileResponse, error) {
	err := i.services.CreateProfile.CreateProfile(ctx, create_profile.NewCreateProfileDTO(
		req.GetUserId(),
		req.GetName(),
		req.GetEmail(),
		req.GetIsTaskAllowed(),
	))
	if err != nil {
		return nil, err
	}

	return &profileV1.CreateProfileResponse{}, nil
}
