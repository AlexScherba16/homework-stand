package check_permission

import (
	"context"

	"profile-service/internal/domain/entity"
)

type ProfileProvider interface {
	GetProfile(ctx context.Context, userID int64) (*entity.Profile, error)
}

type Service struct {
	provider ProfileProvider
}

func NewService(provider ProfileProvider) *Service {
	return &Service{provider: provider}
}

func (s *Service) CheckPermission(ctx context.Context, userID int64) (bool, error) {
	profile, err := s.provider.GetProfile(ctx, userID)
	if err != nil {
		return false, err
	}

	return profile.IsTaskAllowed, nil
}
