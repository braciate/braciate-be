package middleware

import (
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/response"
	"github.com/braciate/braciate-be/internal/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var (
	ErrForbidden = response.NewError(http.StatusForbidden, "forbidden access")
)

func OnlyTrustedRole(roles ...entity.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		account, err := utils.GetUserFromContext(c)
		if err != nil {
			return ErrForbidden
		}

		for _, role := range roles {
			if role == account.Role {
				return c.Next()
			}
		}

		return ErrForbidden
	}
}
