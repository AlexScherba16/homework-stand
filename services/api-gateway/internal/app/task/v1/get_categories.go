package v1

import (
	"context"

	taskV1 "api-gateway/internal/pkg/pb/api-gateway/task/v1"

	externalV1 "api-gateway/internal/pkg/pb/external/task-service/task/v1"

	"github.com/samber/lo"
)

func (i *Implementation) GetCategories(ctx context.Context, req *taskV1.GetCategoriesRequest) (*taskV1.GetCategoriesResponse, error) {
	response, err := i.external.GetCategories(ctx, &externalV1.GetCategoriesRequest{})
	if err != nil {
		return nil, err
	}

	return &taskV1.GetCategoriesResponse{
		Categories: lo.Map(response.GetCategories(), func(category *externalV1.GetCategoriesResponse_Category, _ int) *taskV1.GetCategoriesResponse_Category {
			return &taskV1.GetCategoriesResponse_Category{
				Id:    category.GetId(),
				Title: category.GetTitle(),
				Desc:  category.GetDesc(),
				Price: category.GetPrice(),
			}
		}),
	}, nil
}
