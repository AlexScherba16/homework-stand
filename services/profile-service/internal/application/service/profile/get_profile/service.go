package get_profile

import (
	"context"

	"profile-service/internal/domain/entity"
)

type ProfileProvider interface {
	GetProfile(ctx context.Context, userID int64) (*entity.Profile, error)
}

type TaskProvider interface {
	GetUserTaskCount(ctx context.Context, userID int64) (uint64, error)
}

type Service struct {
	profileProvider ProfileProvider
	taskProvider    TaskProvider
}

func NewService(profileProvider ProfileProvider, taskProvider TaskProvider) *Service {
	return &Service{profileProvider: profileProvider, taskProvider: taskProvider}
}

func (s *Service) GetProfile(ctx context.Context, userID int64) (*entity.Profile, error) {
	// получаем профиль
	profile, err := s.profileProvider.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	// получаем кол-во задач
	taskCount, err := s.taskProvider.GetUserTaskCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	profile.WithTariff(taskCount)

	return profile, nil
}
