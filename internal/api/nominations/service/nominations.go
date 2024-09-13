package nominationsService

import (
	"errors"
	"fmt"
	"time"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
	"golang.org/x/net/context"
)

func (s *NominationsService) CreateNomination(ctx context.Context, nominationReq entity.Nominations) (nominations.NominationResponse, error) {
	nominationRepo, err := s.nominationsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating nomination repository: %v", err)
		return nominations.NominationResponse{}, err
	}

	generateID, err := utils.GenerateRandomString(24)
	if err != nil {
		s.log.Errorf("error generating ID: %v", err)
		return nominations.NominationResponse{}, err
	}

	nominationReq.ID = fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID)

	newNomination, err := nominationRepo.CreateNomination(ctx, nominationReq)
	if err != nil {
		if errors.Is(err, nominations.ErrForeignKeyViolation) {
			s.log.Errorf("Foreign key violation: %v", err)
			return nominations.NominationResponse{}, nominations.ErrForeignKeyViolation
		}
		if errors.Is(err, nominations.ErrUniqueViolation) {
			s.log.Errorf("Unique violation: %v", err)
			return nominations.NominationResponse{}, nominations.ErrUniqueViolation
		}
		s.log.Errorf("CreateNomination err: %v", err)
		return nominations.NominationResponse{}, err
	}

	return nominations.NominationResponse{
		ID:         newNomination.ID,
		Name:       newNomination.Name,
		CategoryID: newNomination.CategoryID,
	}, nil
}

func (s *NominationsService) GetAllNominationsByCategoryID(ctx context.Context, id string) ([]nominations.NominationResponse, error) {
	nominationRepo, err := s.nominationsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating nomination repository: %v", err)
		return nil, err
	}

	getNomination, err := nominationRepo.GetAllNominationsByCategoryID(ctx, id)
	if err != nil {
		s.log.Errorf("GetNomination err: %v", err)
		return nil, err
	}

	var nominationResponses []nominations.NominationResponse
	for _, nomination := range getNomination {
		nominationResponses = append(nominationResponses, nominations.NominationResponse{
			ID:         nomination.ID,
			Name:       nomination.Name,
			CategoryID: nomination.CategoryID,
		})

	}

	return nominationResponses, nil

}
