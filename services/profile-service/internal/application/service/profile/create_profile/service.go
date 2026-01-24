package create_profile

import (
	"context"

	"profile-service/internal/domain/entity"
)

type ProfileCreator interface {
	CreateProfile(ctx context.Context, p *entity.Profile) error
}

type Service struct {
	creator ProfileCreator
}

func NewService(creator ProfileCreator) *Service {
	return &Service{creator: creator}
}

func (s *Service) CreateProfile(ctx context.Context, dto CreateProfileDTO) error {
	profile := entity.NewProfile(dto.userID, dto.name, dto.email)

	profile.AllowTask(dto.isTaskAllowed)

	return s.creator.CreateProfile(ctx, profile)
}
