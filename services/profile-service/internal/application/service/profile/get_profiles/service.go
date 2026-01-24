package get_profiles

import (
	"context"

	"profile-service/internal/domain/entity"
)

type ProfileProvider interface {
	GetProfiles(ctx context.Context, userIDs ...int64) (map[int64]*entity.Profile, error)
}

type Service struct {
	profileProvider ProfileProvider
}

func NewService(profileProvider ProfileProvider) *Service {
	return &Service{profileProvider: profileProvider}
}

func (s *Service) GetProfiles(ctx context.Context, userIDs ...int64) (map[int64]*entity.Profile, error) {
	return s.profileProvider.GetProfiles(ctx, userIDs...)
}
