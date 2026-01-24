package v1

import (
	"context"

	"task-service/internal/domain/entity"
	"task-service/internal/pkg/convert"
	taskV1 "task-service/internal/pkg/pb/task-service/task/v1"

	"github.com/samber/lo"
)

func (i *Implementation) GetCategories(ctx context.Context, req *taskV1.GetCategoriesRequest) (*taskV1.GetCategoriesResponse, error) {
	categories, err := i.services.GetCategories.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	return &taskV1.GetCategoriesResponse{
		Categories: lo.Map(categories, func(category entity.Category, _ int) *taskV1.GetCategoriesResponse_Category {
			return &taskV1.GetCategoriesResponse_Category{
				Id:    category.ID,
				Title: category.Title,
				Price: convert.DecimalToMoney(category.Price),
				Desc:  category.Desc,
			}
		}),
	}, nil
}
