package lkmsService

import (
	lkmsRepository "github.com/braciate/braciate-be/internal/api/lkms/repository"
	"github.com/sirupsen/logrus"
)

type LkmsService struct {
	lkmsRepository lkmsRepository.RepositoryItf
	log            *logrus.Logger
}

type LkmsServiceItf interface {
}

func (s *LkmsService) New(log *logrus.Logger, repo lkmsRepository.RepositoryItf) LkmsServiceItf {
	return &LkmsService{
		lkmsRepository: repo,
		log:            log,
	}
}
