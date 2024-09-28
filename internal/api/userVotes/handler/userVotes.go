package votesHandler

import (
	"context"
	"time"

	"github.com/braciate/braciate-be/internal/api/userVotes"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (h *UserVotesHandler) CreateUserVotesHandler(ctx *fiber.Ctx) error {
	var req userVotes.UserVotesRequest

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := h.parseAndBindUserVotesRequest(ctx)
	if err != nil {
		return err
	}

	votesReq := entity.UserVotes{
		UserID:       req.UserID,
		LkmID:        req.LkmID,
		NominationID: req.NominationID,
	}

	res, err := h.userVotesService.CreateNomination(c, votesReq)
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

func (h *UserVotesHandler) GetAllUserVotesByNomination(ctx *fiber.Ctx) error {
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

	res, err := h.userVotesService.GetAllUserVotesByNomination(c, id)
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

func (h *UserVotesHandler) DeleteUserVotes(ctx *fiber.Ctx) error {
	var id string
	var err error
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id = ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.StatusMessage(fiber.StatusBadRequest))
	}

	res, err := h.userVotesService.DeleteUserVotes(c, id)
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
