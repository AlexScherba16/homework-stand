package v1

import (
	"context"

	profileV1 "profile-service/internal/pkg/pb/profile-service/profile/v1"
)

func (i *Implementation) GetProfileList(ctx context.Context, req *profileV1.GetProfileListRequest) (*profileV1.GetProfileListResponse, error) {
	profiles, err := i.services.GetProfiles.GetProfiles(ctx, req.GetUserIds()...)
	if err != nil {
		return nil, err
	}

	resp := &profileV1.GetProfileListResponse{
		Profiles: make([]*profileV1.GetProfileListResponse_Profile, 0, len(profiles)),
	}

	for _, profile := range profiles {
		resp.Profiles = append(resp.Profiles, &profileV1.GetProfileListResponse_Profile{
			UserId:        profile.UserID,
			Name:          profile.Name,
			Email:         profile.Email,
			IsTaskAllowed: profile.IsTaskAllowed,
		})
	}
	return resp, nil
}
