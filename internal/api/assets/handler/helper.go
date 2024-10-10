package assetsHandler

import (
	"github.com/braciate/braciate-be/internal/api/assets"
	"github.com/gofiber/fiber/v2"
)

func (h *AssetsHandler) parseAndBindAssetsRequest(ctx *fiber.Ctx) (assets.AssetsRequest, error) {
	var req assets.AssetsRequest
	if err := ctx.BodyParser(&req); err != nil {
		return req, err
	}

	if err := h.validate.Struct(req); err != nil {
		return req, err
	}

	return req, nil
}
