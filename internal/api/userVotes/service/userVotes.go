package votesService

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/braciate/braciate-be/internal/api/userVotes"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
)

func (s *UserVotesService) CreateNomination(ctx context.Context, votesReq entity.UserVotes) (userVotes.UserVotesResponse, error) {
	userVotesRepo, err := s.UserVotesRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating nomination repository: %v", err)
		return userVotes.UserVotesResponse{}, err
	}

	generateID, err := utils.GenerateRandomString(24)
	if err != nil {
		s.log.Errorf("error generating ID: %v", err)
		return userVotes.UserVotesResponse{}, err
	}

	votesReq.ID = fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID)

	newVotes, err := userVotesRepo.CreateUserVotes(ctx, votesReq)
	if err != nil {
		if errors.Is(err, userVotes.ErrForeignKeyViolation) {
			s.log.Errorf("Foreign key violation: %v", err)
			return userVotes.UserVotesResponse{}, userVotes.ErrForeignKeyViolation
		}
		if errors.Is(err, userVotes.ErrUniqueViolation) {
			s.log.Errorf("Unique violation: %v", err)
			return userVotes.UserVotesResponse{}, userVotes.ErrUniqueViolation
		}
		s.log.Errorf("CreateUserVotes err: %v", err)
		return userVotes.UserVotesResponse{}, err
	}

	return userVotes.UserVotesResponse{
		ID:           newVotes.ID,
		UserID:       newVotes.UserID,
		LkmID:        newVotes.LkmID,
		NominationID: newVotes.NominationID,
	}, nil
}
