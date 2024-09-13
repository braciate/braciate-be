package nominationsService

import (
	"fmt"
	"time"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
	"golang.org/x/net/context"
)

func (s *NominationsService) CreateCategory(ctx context.Context, categoyReq entity.Categories) (nominations.CategoryResponse, error) {
	categoryRepo, err := s.nominationsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating category repository: %v", err)
		return nominations.CategoryResponse{}, err
	}

	generateID, err := utils.GenerateRandomString(24)
	if err != nil {
		s.log.Errorf("error generating ID: %v", err)
		return nominations.CategoryResponse{}, err
	}

	categoyReq.ID = fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID)

	newCategory, err := categoryRepo.CreateCategory(ctx, categoyReq)
	if err != nil {
		return nominations.CategoryResponse{}, err
	}

	return nominations.CategoryResponse{
		ID:   newCategory.ID,
		Name: newCategory.Name,
	}, nil
}
func (s *NominationsService) GetAllCategories(ctx context.Context) ([]nominations.CategoryResponse, error) {
	categoriesRepo, err := s.nominationsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating categories repository: %v", err)
		return nil, err
	}

	getCategories, err := categoriesRepo.GetAllCategories(ctx)
	if err != nil {
		s.log.Errorf("GetCategories err : %v", err)
		return nil, err
	}

	var categoryResponses []nominations.CategoryResponse
	for _, category := range getCategories {
		categoryResponses = append(categoryResponses, nominations.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return categoryResponses, nil
}

func (s *NominationsService) UpdateCategory(ctx context.Context, req entity.Categories, id string) (nominations.CategoryResponse, error) {
	categoryRepo, err := s.nominationsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating category repository: %v", err)
		return nominations.CategoryResponse{}, err
	}

	oldCategory, err := categoryRepo.GetCategoryByID(ctx, id)
	if err != nil {
		s.log.Errorf("error getting category by ID: %v", err)
		return nominations.CategoryResponse{}, err
	}

	updatedCategory := oldCategory
	if oldCategory.Name != req.Name {
		updatedCategory.Name = req.Name
	} else {
		return nominations.CategoryResponse{
			ID:   updatedCategory.ID,
			Name: updatedCategory.Name,
		}, err
	}

	updatedCategory, err = categoryRepo.UpdateCategory(ctx, updatedCategory)
	if err != nil {
		s.log.Errorf("error updating category: %v", err)
		return nominations.CategoryResponse{}, err
	}

	return nominations.CategoryResponse{
		ID:   updatedCategory.ID,
		Name: updatedCategory.Name,
	}, nil
}

func (s *NominationsService) DeleteCategory(ctx context.Context, id string) (nominations.CategoryResponse, error) {
	categoryRepo, err := s.nominationsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating nomination repository: %v", err)
		return nominations.CategoryResponse{}, err
	}

	deleted, err := categoryRepo.DeleteCategory(ctx, id)
	if err != nil {
		s.log.Errorf("GetNomination err: %v", err)
		return nominations.CategoryResponse{}, err
	}

	res := nominations.CategoryResponse{
		ID:   deleted.ID,
		Name: deleted.Name,
	}

	return res, nil

}
