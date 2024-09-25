package lkmsService

import (
	"context"
	"fmt"
	"time"

	"github.com/braciate/braciate-be/internal/api/lkms"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
)

func (s *LkmsService) CreateLkm(ctx context.Context, req entity.Lkms) (lkms.LkmsResponse, error) {
	lkmRepo, err := s.lkmsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating lkm category: %v", err)
		return lkms.LkmsResponse{}, err
	}

	generateID, err := utils.GenerateRandomString(24)
	if err != nil {
		s.log.Errorf("error generating ID: %v", err)
		s.log.Errorf("error generating ID: %v", err)
		return lkms.LkmsResponse{}, err
	}

	req.ID = fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID)

	newLkm, err := lkmRepo.CreateLkms(ctx, req)
	if err != nil {
		return lkms.LkmsResponse{}, err
	}

	return lkms.LkmsResponse{
		ID:         newLkm.ID,
		Name:       newLkm.Name,
		CategoryID: newLkm.CategoryID,
		Type:       newLkm.Type,
		LogoLink:   newLkm.LogoLink,
	}, err
}
