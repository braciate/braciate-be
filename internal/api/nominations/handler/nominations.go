package nominationsHandler

import (
	"time"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"golang.org/x/net/context"
)

func (h *NominationHandler) CreateNominationHandler(ctx *fiber.Ctx) error {
	var req nominations.NominationRequest

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := h.parseAndBindNominationsRequest(ctx)
	if err != nil {
		return err
	}

	nominationReq := entity.Nominations{
		Name:       req.Name,
		CategoryID: req.CategoryID,
	}

	res, err := h.nominationsService.CreateNomination(c, nominationReq)
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

func (h *NominationHandler) GetAllNominatonsByCategoryID(ctx *fiber.Ctx) error {
	var id string
	var err error
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id = ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.StatusMessage(fiber.StatusBadRequest))
	}

	res, err := h.nominationsService.GetAllNominationsByCategoryID(c, id)
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

func (h *NominationHandler) UpdateNomination(ctx *fiber.Ctx) error {
	var req nominations.NominationRequest

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := h.parseAndBindNominationsRequest(ctx)
	if err != nil {
		return err
	}

	id := ctx.Params("id")

	nominationUpdateReq := entity.Nominations{Name: req.Name, CategoryID: req.CategoryID}

	res, err := h.nominationsService.UpdateNomination(c, nominationUpdateReq, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update nomination",
		})
	}

	select {
	case <-c.Done():
		return ctx.Status(fiber.StatusRequestTimeout).JSON(utils.StatusMessage(fiber.StatusRequestTimeout))
	default:
		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

func (h *NominationHandler) DeleteNomination(ctx *fiber.Ctx) error {
	var id string
	var err error
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id = ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.StatusMessage(fiber.StatusBadRequest))
	}

	res, err := h.nominationsService.DeleteNomination(c, id)
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
