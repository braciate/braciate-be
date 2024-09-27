package votesHandler

import (
	"github.com/braciate/braciate-be/internal/api/userVotes"
	"github.com/gofiber/fiber/v2"
)

func (h *UserVotesHandler) parseAndBindUserVotesRequest(ctx *fiber.Ctx) (userVotes.UserVotesRequest, error) {
	var req userVotes.UserVotesRequest
	if err := ctx.BodyParser(&req); err != nil {
		return req, err
	}

	if err := h.validate.Struct(req); err != nil {
		return req, err
	}

	return req, nil
}
