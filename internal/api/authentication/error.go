package authentication

import (
	"github.com/braciate/braciate-be/internal/pkg/response"
	"net/http"
)

var (
	ErrIntializeSomething = response.NewError(http.StatusBadRequest, "something went wrong")
)
