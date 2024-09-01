package categoriesService

import (
	"fmt"
	"time"

	"github.com/braciate/braciate-be/internal/api/categories"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
	"golang.org/x/net/context"
)

func (s *CategoriesService) CreateCategory(ctx context.Context, request categories.CreateCategoryRequest) (categories.CreateCategoryResponse, error) {
	categoryRepo, err := s.categoriesRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating category repository: %v", err)
		return categories.CreateCategoryResponse{}, err
	}

	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	generateID, err := utils.GenerateRandomString(24)
	if err != nil {
		s.log.Errorf("error generating ID: %v", err)
		return categories.CreateCategoryResponse{}, err
	}

	categoryReq := entity.Categories{
		ID:   fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID),
		Name: request.Name,
	}

	_, err = categoryRepo.CreateCategory(c, categoryReq)
	if err != nil {
		return categories.CreateCategoryResponse{}, err
	}

	return categories.CreateCategoryResponse{
		ID:   categoryReq.ID,
		Name: categoryReq.Name,
	}, nil
}
