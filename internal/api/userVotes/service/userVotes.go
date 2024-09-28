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
		s.log.Errorf("error creating user votes repository: %v", err)
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

func (s *UserVotesService) GetAllUserVotesByNomination(ctx context.Context, id string) ([]userVotes.UserVotesResponse, error) {
	userVotesRepo, err := s.UserVotesRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating user votes repository: %v", err)
		return nil, err
	}

	getVotes, err := userVotesRepo.GetAllUserVotesByNomination(ctx, id)
	if err != nil {
		s.log.Errorf("Get user votes err: %v", err)
		return nil, err
	}

	var voteResponses []userVotes.UserVotesResponse
	for _, vote := range getVotes {
		voteResponses = append(voteResponses, userVotes.UserVotesResponse{
			ID:           vote.ID,
			UserID:       vote.UserID,
			LkmID:        vote.LkmID,
			NominationID: vote.NominationID,
		})

	}

	return voteResponses, nil

}

func (s *UserVotesService) DeleteUserVotes(ctx context.Context, id string) (userVotes.UserVotesResponse, error) {
	userVotesRepo, err := s.UserVotesRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating user votes repository: %v", err)
		return userVotes.UserVotesResponse{}, err
	}

	deleted, err := userVotesRepo.DeleteUserVotes(ctx, id)
	if err != nil {
		s.log.Errorf("DeleteUserVotes err: %v", err)
		return userVotes.UserVotesResponse{}, err
	}

	res := userVotes.UserVotesResponse{
		ID:           deleted.ID,
		UserID:       deleted.UserID,
		NominationID: deleted.NominationID,
		LkmID:        deleted.LkmID,
	}

	return res, nil

}
