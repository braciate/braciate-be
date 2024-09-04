package nominationsHandler

import (
	"errors"
	"time"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"golang.org/x/net/context"
)

func (h *NominationHandler) CreateNominationHandler(ctx *fiber.Ctx) error {
	var req nominations.CreateNominationRequest

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
		if errors.Is(err, nominations.ErrForeignKeyViolation) {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				utils.StatusMessage(fiber.StatusBadRequest))
		}

		if errors.Is(err, nominations.ErrUniqueViolation) {
			return ctx.Status(fiber.StatusConflict).JSON(
				utils.StatusMessage(fiber.StatusConflict))
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.StatusMessage(fiber.StatusInternalServerError))
	}

	select {
	case <-c.Done():
		return ctx.Status(fiber.StatusRequestTimeout).JSON(
			utils.StatusMessage(fiber.StatusRequestTimeout))
	default:
		return ctx.Status(fiber.StatusOK).JSON(res)
	}
}

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
