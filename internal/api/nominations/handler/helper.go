package nominationsHandler

import (
	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/gofiber/fiber/v2"
)

func (h *NominationHandler) parseAndBindNominationsRequest(ctx *fiber.Ctx) (nominations.CreateNominationRequest, error) {
	var req nominations.CreateNominationRequest
	if err := ctx.BodyParser(&req); err != nil {
		return req, err
	}

	if err := h.validate.Struct(req); err != nil {
		return req, err
	}

	return req, nil
}

func (h *NominationHandler) parseAndBindCategoriesRequest(ctx *fiber.Ctx) (nominations.CreateCategoryRequest, error) {
	var req nominations.CreateCategoryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return req, err
	}

	if err := h.validate.Struct(req); err != nil {
		return req, err
	}

	return req, nil
}
