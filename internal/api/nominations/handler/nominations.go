package nominationsHandler

import (
	"time"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"golang.org/x/net/context"
)

func (h *NominationHandler) CreateNominationHandler(ctx *fiber.Ctx) error {
	var req nominations.CreateNominationRequest
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := h.parseAndBindRequest(ctx)
	if err != nil {
		return err
	}

	res, err := h.nominationsService.CreateNomination(c, req)
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
