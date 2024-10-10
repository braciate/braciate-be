package assetsHandler

import (
	"context"
	"time"

	"github.com/braciate/braciate-be/internal/api/assets"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (h *AssetsHandler) CreateAssetsHandler(ctx *fiber.Ctx) error {
	var req assets.AssetsRequest

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := h.parseAndBindAssetsRequest(ctx)
	if err != nil {
		return err
	}

	votesReq := entity.Assets{
		UserID:       req.UserID,
		LkmID:        req.LkmID,
		NominationID: req.NominationID,
	}

	res, err := h.assetsService.CreateAssets(c, votesReq)
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

func (h *AssetsHandler) GetAllAssetsByNomination(ctx *fiber.Ctx) error {
	var (
		id  string
		err error
	)
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id = ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.StatusMessage(fiber.StatusBadRequest))
	}

	res, err := h.assetsService.GetAllAssetsByNomination(c, id)
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

func (h *AssetsHandler) DeleteAssets(ctx *fiber.Ctx) error {
	var id string
	var err error
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id = ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.StatusMessage(fiber.StatusBadRequest))
	}

	res, err := h.assetsService.DeleteAssets(c, id)
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
