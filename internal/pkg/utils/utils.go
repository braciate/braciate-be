package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/braciate/braciate-be/internal/api/authentication"
	"github.com/gofiber/fiber/v2"
)

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

func GetUserFromContext(ctx *fiber.Ctx) (authentication.UserClaims, error) {
	userCtx := ctx.Locals("user")
	if userCtx == nil {
		return authentication.UserClaims{}, fmt.Errorf("user not found in context")
	}

	return userCtx.(authentication.UserClaims), nil
}
