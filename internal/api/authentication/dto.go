package authentication

import "github.com/braciate/braciate-be/internal/entity"

type SigninRequest struct {
	NimEmail string `json:"nim_email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SigninResponse struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   int64  `json:"expired_at"`
}

type UserClaims struct {
	Username string `json:"username"`
	Role     entity.UserRole
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nim      string `json:"nim"`
}
