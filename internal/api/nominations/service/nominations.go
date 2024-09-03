package nominationsService

import (
	"fmt"
	"time"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
	"golang.org/x/net/context"
)

func (s *NominationsService) CreateNomination(ctx context.Context, request nominations.CreateNominationRequest) (nominations.CreateNominationResponse, error) {
	nominationRepo, err := s.nominationsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating nomination repository: %v", err)
		return nominations.CreateNominationResponse{}, err
	}

	generateID, err := utils.GenerateRandomString(24)
	if err != nil {
		s.log.Errorf("error generating ID: %v", err)
		return nominations.CreateNominationResponse{}, err
	}

	nominationReq := entity.Nominations{
		ID:         fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID),
		Name:       request.Name,
		CategoryID: request.CategoryID,
	}

	_, err = nominationRepo.CreateNomination(ctx, nominationReq)
	if err != nil {
		return nominations.CreateNominationResponse{}, err
	}

	return nominations.CreateNominationResponse{
		ID:         nominationReq.ID,
		Name:       nominationReq.Name,
		CategoryID: nominationReq.CategoryID,
	}, nil
}
