package lkmsHandler

import (
	"context"
	"strconv"
	"time"

	"github.com/braciate/braciate-be/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (h *LkmsHandler) CreateLkms(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	name := ctx.FormValue("name")

	logo, err := ctx.FormFile("logo")

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get 'logo' from form data",
		})
		return err
	}

	categoryID := ctx.FormValue("category_id")

	typeLkm := ctx.FormValue("type")

	typeInt, err := strconv.Atoi(typeLkm)
	if err != nil {
		return err
	}

	lkmsReq := entity.Lkms{
		Name:       name,
		CategoryID: categoryID,
		Type:       typeInt,
	}

	res, err := h.lkmsService.CreateLkm(c, lkmsReq, logo)
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
