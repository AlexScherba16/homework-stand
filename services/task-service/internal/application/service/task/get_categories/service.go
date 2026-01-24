package get_categories

import (
	"context"

	"task-service/internal/domain/entity"
)

type Category interface {
	GetAllCategories(_ context.Context) ([]entity.Category, error)
}

type Service struct {
	category Category
}

func NewService(category Category) *Service {
	return &Service{category: category}
}

func (s *Service) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return s.category.GetAllCategories(ctx)
}
