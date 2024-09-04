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
