package nominationsService

import (
	"fmt"
	"time"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
	"golang.org/x/net/context"
)

func (s *NominationsService) CreateNomination(ctx context.Context, nominationReq entity.Nominations) (nominations.CreateNominationResponse, error) {
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

	nominationReq.ID = fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID)

	newNomination, err := nominationRepo.CreateNomination(ctx, nominationReq)
	if err != nil {
		return nominations.CreateNominationResponse{}, err
	}

	return nominations.CreateNominationResponse{
		ID:         newNomination.ID,
		Name:       newNomination.Name,
		CategoryID: newNomination.CategoryID,
	}, nil
}
