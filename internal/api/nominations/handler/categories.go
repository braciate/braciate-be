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
	var req nominations.CategoryRequest
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

func (h *NominationHandler) GetAllCategories(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.nominationsService.GetAllCategories(c)
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

func (h *NominationHandler) UpdateCategory(ctx *fiber.Ctx) error {
	var req nominations.CategoryRequest

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := h.parseAndBindCategoriesRequest(ctx)
	if err != nil {
		return err
	}

	id := ctx.Params("id")

	categoryUpdateReq := entity.Categories{Name: req.Name}

	res, err := h.nominationsService.UpdateCategory(c, categoryUpdateReq, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update category",
		})
	}

	select {
	case <-c.Done():
		return ctx.Status(fiber.StatusRequestTimeout).JSON(utils.StatusMessage(fiber.StatusRequestTimeout))
	default:
		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

func (h *NominationHandler) DeleteCategory(ctx *fiber.Ctx) error {
	var id string
	var err error
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id = ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.StatusMessage(fiber.StatusBadRequest))
	}

	res, err := h.nominationsService.DeleteCategory(c, id)
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
