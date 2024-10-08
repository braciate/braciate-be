package assetsService

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/braciate/braciate-be/internal/api/assets"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
)

func (s *AssetsService) CreateAssets(ctx context.Context, req entity.Assets) (assets.AssetsResponse, error) {
	AssetsRepo, err := s.AssetsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating user votes repository: %v", err)
		return assets.AssetsResponse{}, err
	}

	generateID, err := utils.GenerateRandomString(24)
	if err != nil {
		s.log.Errorf("error generating ID: %v", err)
		return assets.AssetsResponse{}, err
	}

	req.ID = fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID)

	newVotes, err := AssetsRepo.CreateAssets(ctx, req)
	if err != nil {
		if errors.Is(err, assets.ErrForeignKeyViolation) {
			s.log.Errorf("Foreign key violation: %v", err)
			return assets.AssetsResponse{}, assets.ErrForeignKeyViolation
		}
		if errors.Is(err, assets.ErrUniqueViolation) {
			s.log.Errorf("Unique violation: %v", err)
			return assets.AssetsResponse{}, assets.ErrUniqueViolation
		}
		s.log.Errorf("CreateAssets err: %v", err)
		return assets.AssetsResponse{}, err
	}

	return assets.AssetsResponse{
		ID:           newVotes.ID,
		UserID:       newVotes.UserID,
		LkmID:        newVotes.LkmID,
		NominationID: newVotes.NominationID,
	}, nil
}

func (s *AssetsService) GetAllAssetsByNomination(ctx context.Context, id string) ([]assets.AssetsResponse, error) {
	AssetsRepo, err := s.AssetsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating user votes repository: %v", err)
		return nil, err
	}

	getAsset, err := AssetsRepo.GetAllAssetsByNomination(ctx, id)
	if err != nil {
		s.log.Errorf("Get user votes err: %v", err)
		return nil, err
	}

	var voteResponses []assets.AssetsResponse
	for _, vote := range getAsset {
		voteResponses = append(voteResponses, assets.AssetsResponse{
			ID:           vote.ID,
			UserID:       vote.UserID,
			LkmID:        vote.LkmID,
			NominationID: vote.NominationID,
		})

	}

	return voteResponses, nil

}

func (s *AssetsService) DeleteAssets(ctx context.Context, id string) (assets.AssetsResponse, error) {
	AssetsRepo, err := s.AssetsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating user votes repository: %v", err)
		return assets.AssetsResponse{}, err
	}

	deleted, err := AssetsRepo.DeleteAssets(ctx, id)
	if err != nil {
		s.log.Errorf("DeleteAssets err: %v", err)
		return assets.AssetsResponse{}, err
	}

	res := assets.AssetsResponse{
		ID:           deleted.ID,
		UserID:       deleted.UserID,
		NominationID: deleted.NominationID,
		LkmID:        deleted.LkmID,
	}

	return res, nil

}
