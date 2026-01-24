package v1

import (
	"context"

	profileV1 "api-gateway/internal/pkg/pb/api-gateway/profile/v1"
	externalV1 "api-gateway/internal/pkg/pb/external/profile-service/profile/v1"
)

func (i *Implementation) GetProfile(ctx context.Context, req *profileV1.GetProfileRequest) (*profileV1.GetProfileResponse, error) {
	response, err := i.external.GetProfile(ctx, &externalV1.GetProfileRequest{
		UserId: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}

	return &profileV1.GetProfileResponse{
		UserId: response.GetUserId(),
		Name:   response.GetName(),
		Email:  response.GetEmail(),
		Tariff: &profileV1.GetProfileResponse_Tariff{
			Tariff: profileV1.GetProfileResponse_Tariff_Enum(response.GetTariff().GetTariff()),
		},
	}, nil
}
