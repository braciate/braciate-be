package nominationsHandler

import (
	"time"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"golang.org/x/net/context"
)

func (h *NominationHandler) CreateCategoryHandler(ctx *fiber.Ctx) error {
	var req nominations.CreateCategoryRequest
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := h.parseAndBindCategoriesRequest(ctx)
	if err != nil {
		return err
	}

	categoyReq := entity.Categories{Name: req.Name}

	res, err := h.nominationsService.CreateCategory(c, categoyReq)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		return ctx.Status(fiber.StatusRequestTimeout).JSON(
			utils.StatusMessage(fiber.StatusRequestTimeout))
	default:
		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}
