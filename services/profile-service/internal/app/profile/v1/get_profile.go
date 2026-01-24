package v1

import (
	"context"

	profileV1 "profile-service/internal/pkg/pb/profile-service/profile/v1"
)

func (i *Implementation) GetProfile(ctx context.Context, req *profileV1.GetProfileRequest) (*profileV1.GetProfileResponse, error) {
	profile, err := i.services.GetProfile.GetProfile(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &profileV1.GetProfileResponse{
		UserId: profile.UserID,
		Name:   profile.Name,
		Email:  profile.Email,
		Tariff: &profileV1.GetProfileResponse_Tariff{
			Tariff: profileV1.GetProfileResponse_Tariff_Enum(profile.Tariff),
		},
	}, nil
}
