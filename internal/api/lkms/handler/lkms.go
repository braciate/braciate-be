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

func (h *LkmsHandler) GetLkmsByCategoryIDAndType(ctx *fiber.Ctx) error {
	var (
		id, lkmType string
		err         error
	)
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id = ctx.Params("id")
	lkmType = ctx.Params("type")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.StatusMessage(fiber.StatusBadRequest))
	}

	res, err := h.lkmsService.GetLkmsByCategoryIDAndType(c, id, lkmType)
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

func (h *LkmsHandler) UpdateLkms(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := ctx.FormValue("id")

	name := ctx.FormValue("name")

	newLogo, err := ctx.FormFile("logo")
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
		ID:         id,
		Name:       name,
		CategoryID: categoryID,
		Type:       typeInt,
	}

	res, err := h.lkmsService.UpdateLkms(c, lkmsReq, newLogo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update lkm",
		})
	}

	select {
	case <-c.Done():
		return ctx.Status(fiber.StatusRequestTimeout).JSON(utils.StatusMessage(fiber.StatusRequestTimeout))
	default:
		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

func (h *LkmsHandler) DeleteLkm(ctx *fiber.Ctx) error {
	var id string
	var err error
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id = ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.StatusMessage(fiber.StatusBadRequest))
	}

	res, err := h.lkmsService.DeleteLkm(c, id)
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
