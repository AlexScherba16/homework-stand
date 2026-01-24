package v1

import (
	"context"

	profileV1 "api-gateway/internal/pkg/pb/api-gateway/profile/v1"
	externalV1 "api-gateway/internal/pkg/pb/external/profile-service/profile/v1"

	"github.com/samber/lo"
)

func (i *Implementation) GetProfileList(ctx context.Context, req *profileV1.GetProfileListRequest) (*profileV1.GetProfileListResponse, error) {
	response, err := i.external.GetProfileList(ctx, &externalV1.GetProfileListRequest{
		UserIds: req.GetUserIds(),
	})
	if err != nil {
		return nil, err
	}

	return &profileV1.GetProfileListResponse{
		Profiles: lo.Map(response.GetProfiles(), func(profile *externalV1.GetProfileListResponse_Profile, _ int) *profileV1.GetProfileListResponse_Profile {
			return &profileV1.GetProfileListResponse_Profile{
				UserId:        profile.GetUserId(),
				Name:          profile.GetName(),
				Email:         profile.GetEmail(),
				IsTaskAllowed: profile.GetIsTaskAllowed(),
			}
		}),
	}, nil
}
