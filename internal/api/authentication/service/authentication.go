package authService

import (
	"errors"
	"fmt"
	"github.com/braciate/braciate-be/internal/api/authentication"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/bcrypt"
	"github.com/braciate/braciate-be/internal/pkg/utils"
	"golang.org/x/net/context"
	"time"
)

func (s *AuthService) Signin(ctx context.Context, request authentication.SigninRequest) (authentication.SigninResponse, error) {
	authRepo, err := s.authRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating auth repository: %v", err)
		return authentication.SigninResponse{}, authentication.ErrInitializeAuthRepository
	}

	user, err := authRepo.GetUserByEmailOrNIM(ctx, request.NimEmail)
	if err != nil {
		if errors.Is(err, authentication.ErrRecordNotFound) {
			iamUser, err := s.broneAuth.Authenticate(request.NimEmail, request.Password)
			if err != nil {
				s.log.Errorf("error authenticating user: %v", err)
				return authentication.SigninResponse{}, err
			}

			generateID, err := utils.GenerateRandomString(24)
			if err != nil {
				s.log.Errorf("error generating ID: %v", err)
				return authentication.SigninResponse{}, err
			}

			user.ID = fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID)
			user.Username = iamUser.Username
			user.Email = iamUser.Email
			user.StudyProgram = iamUser.StudyProgram
			user.NIM = iamUser.NIM
			user.Faculty = iamUser.Faculty
			user.Role = entity.UserRoleStudent
			hashPass, err := bcrypt.HashPassword(fmt.Sprintf("%s", generateID))
			if err != nil {
				s.log.Errorf("error hashing password: %v", err)
				return authentication.SigninResponse{}, err
			}
			user.Password = hashPass

			_, err = authRepo.CreateUser(ctx, user)
			if err != nil {
				s.log.Errorf("error creating user: %v", err)
				return authentication.SigninResponse{}, err
			}

			accessToken, err := s.generateAccessToken(user)
			if err != nil {
				s.log.Errorf("error generating access token: %v", err)
				return authentication.SigninResponse{}, err
			}

			return authentication.SigninResponse{
				AccessToken: accessToken,
				ExpiredAt:   time.Now().Add(3 * time.Hour).UnixNano(),
			}, nil
		} else {
			s.log.Errorf("error getting user: %v", err)
			return authentication.SigninResponse{}, err
		}
	}

	switch user.Role {
	case entity.UserRoleDelegation:
		if err := bcrypt.ComparePassword(user.Password, request.Password); err != nil {
			s.log.Errorf("error comparing password: %v", err)
			return authentication.SigninResponse{}, authentication.ErrInvalidNimEmailOrPassword
		}
	default:
		iamUser, err := s.broneAuth.Authenticate(request.NimEmail, request.Password)
		if err != nil {
			s.log.Errorf("error authenticating user: %v", err)
			return authentication.SigninResponse{}, err
		}

		user.Username = iamUser.Username
		user.Email = iamUser.Email
		user.StudyProgram = iamUser.StudyProgram
		user.NIM = iamUser.NIM
		user.Faculty = iamUser.Faculty
	}

	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		s.log.Errorf("error generating access token: %v", err)
		return authentication.SigninResponse{}, err
	}

	return authentication.SigninResponse{
		AccessToken: accessToken,
		ExpiredAt:   time.Now().Add(3 * time.Hour).UnixNano(),
	}, nil
}
