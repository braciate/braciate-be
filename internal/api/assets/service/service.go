package assetsService

import (
	"context"

	"github.com/braciate/braciate-be/internal/api/assets"
	assetsRepository "github.com/braciate/braciate-be/internal/api/assets/repository"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/sirupsen/logrus"
)

type AssetsService struct {
	AssetsRepository assetsRepository.RepositoryItf
	log              *logrus.Logger
}

type AssetsServiceItf interface {
	CreateAssets(ctx context.Context, votesReq entity.Assets) (assets.AssetsResponse, error)
	GetAllAssetsByNomination(ctx context.Context, id string) ([]assets.AssetsResponse, error)
	DeleteAssets(ctx context.Context, id string) (assets.AssetsResponse, error)
}

func New(log *logrus.Logger, repo assetsRepository.RepositoryItf) AssetsServiceItf {
	return &AssetsService{
		AssetsRepository: repo,
		log:              log,
	}
}
