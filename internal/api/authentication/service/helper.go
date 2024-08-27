package authService

import (
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/jwt"
	"time"
)

var (
	JWTSecretEnvKey = "JWT_SECRET_ACCESS_TOKEN"
)

func (s *AuthService) generateAccessToken(user entity.User) (string, error) {
	claims := map[string]interface{}{
		"username": user.Username,
		"id":       user.ID,
		"role":     user.Role,
		"email":    user.Email,
		"nim":      user.NIM,
	}
	return jwt.Sign(claims, JWTSecretEnvKey, 3*time.Hour)
}
