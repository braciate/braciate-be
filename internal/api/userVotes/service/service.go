package votesService

import (
	"context"

	uservotes "github.com/braciate/braciate-be/internal/api/userVotes"
	votesRepository "github.com/braciate/braciate-be/internal/api/userVotes/repository"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/sirupsen/logrus"
)

type UserVotesService struct {
	UserVotesRepository votesRepository.RepositoryItf
	log                 *logrus.Logger
}

type UserVotesServiceItf interface {
	CreateNomination(ctx context.Context, votesReq entity.UserVotes) (uservotes.UserVotesResponse, error)
}

func New(log *logrus.Logger, repo votesRepository.RepositoryItf) UserVotesServiceItf {
	return &UserVotesService{
		UserVotesRepository: repo,
		log:                 log,
	}
}
