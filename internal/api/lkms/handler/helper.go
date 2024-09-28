package lkmsHandler

import (
	"github.com/braciate/braciate-be/internal/api/lkms"
	"github.com/gofiber/fiber/v2"
)

func (h *LkmsHandler) parseAndBindLkmsRequest(ctx *fiber.Ctx) (lkms.LkmsRequest, error) {
	var req lkms.LkmsRequest
	if err := ctx.BodyParser(&req); err != nil {
		return req, err
	}

	if err := h.validate.Struct(req); err != nil {
		return req, err
	}

	return req, nil
}
