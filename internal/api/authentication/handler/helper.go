package authHandler

import (
	"github.com/braciate/braciate-be/internal/api/authentication"
	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) parseAndBindRequest(ctx *fiber.Ctx) (authentication.SigninRequest, error) {
	var req authentication.SigninRequest
	if err := ctx.BodyParser(&req); err != nil {
		return req, err
	}

	if err := h.validate.Struct(req); err != nil {
		return req, err
	}

	return req, nil
}
