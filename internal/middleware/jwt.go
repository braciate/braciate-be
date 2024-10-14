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
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return ErrUnauthorized
		}

		// Safely retrieve and type-assert claims
		userID, ok := claims["account_id"].(string)
		if !ok {
			return ErrUnauthorized // or handle the error accordingly
		}

		nim, ok := claims["nim"].(string)
		if !ok {
			nim = "" // Default to empty string if missing
		}

		role, ok := claims["role"].(float64) // roles might be stored as float64
		if !ok {
			return ErrUnauthorized // handle role extraction issue
		}

		email, ok := claims["email"].(string)
		if !ok {
			email = "" // Default to empty string if missing
		}

		username, ok := claims["username"].(string)
		if !ok {
			username = "" // Default to empty string if missing
		}

		// Set user claims in context for the next handler
		user := authentication.UserClaims{
			ID:       userID,
			Nim:      nim,
			Role:     entity.UserRole(role),
			Email:    email,
			Username: username,
		}
		c.Locals("user", user)
		return c.Next()
	}
}
