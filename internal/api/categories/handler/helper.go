package categoriesHandler

import (
	"github.com/braciate/braciate-be/internal/api/categories"
	"github.com/gofiber/fiber/v2"
)

func (h *CategoryHandler) parseAndBindRequest(ctx *fiber.Ctx) (categories.CreateCategoryRequest, error) {
	var req categories.CreateCategoryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return req, err
	}

	if err := h.validate.Struct(req); err != nil {
		return req, err
	}

	return req, nil
}
