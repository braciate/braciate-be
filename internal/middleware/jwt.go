package middleware

import (
	"github.com/braciate/braciate-be/internal/api/authentication"
	"github.com/braciate/braciate-be/internal/entity"
	jwtHelper "github.com/braciate/braciate-be/internal/pkg/jwt"
	"github.com/braciate/braciate-be/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

var (
	ErrUnauthorized = response.NewError(http.StatusUnauthorized, "unauthorized, access token invalid or expired")
)

const (
	AccessTokenSecret = "JWT_ACCESS_TOKEN_SECRET"
)

func JWTAccessToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("Authorization") == "" {
			return ErrUnauthorized
		}

		if !strings.Contains(c.Get("Authorization"), "Bearer") {
			return ErrUnauthorized
		}

		token, err := jwtHelper.VerifyTokenHeader(c, AccessTokenSecret)
		if err != nil {
			return ErrUnauthorized
		} else {
			claims := token.Claims.(jwt.MapClaims)
			user := authentication.UserClaims{
				ID:       claims["account_id"].(string),
				Nim:      claims["nim"].(string),
				Role:     entity.UserRole(claims["role"].(float64)),
				Email:    claims["email"].(string),
				Username: claims["username"].(string),
			}
			c.Locals("user", user)
			return c.Next()
		}
	}
}
