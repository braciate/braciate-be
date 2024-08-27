package authService

import (
	"github.com/braciate/braciate-be/internal/api/authentication"
	authRepository "github.com/braciate/braciate-be/internal/api/authentication/repository"
	broneAuth "github.com/braciate/braciate-be/internal/pkg/brone_auth"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type AuthService struct {
	authRepository authRepository.RepositoryItf
	broneAuth      broneAuth.BroneAuth
	log            *logrus.Logger
}

type AuthServiceItf interface {
	Signin(ctx context.Context, request authentication.SigninRequest) (authentication.SigninResponse, error)
}

func New(log *logrus.Logger, repo authRepository.RepositoryItf, broneAuth broneAuth.BroneAuth) AuthServiceItf {
	return &AuthService{
		authRepository: repo,
		broneAuth:      broneAuth,
		log:            log,
	}
}
