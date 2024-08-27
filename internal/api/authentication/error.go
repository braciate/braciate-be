package authentication

import (
	"github.com/braciate/braciate-be/internal/pkg/response"
	"net/http"
)

// Error from repository layer
var (
	ErrCommitTransaction   = response.NewError(http.StatusInternalServerError, "failed to commit transaction")
	ErrRollbackTransaction = response.NewError(http.StatusInternalServerError, "failed to rollback transaction")
	ErrRecordNotFound      = response.NewError(http.StatusNotFound, "record not found")
)

// Error from service layer
var (
	ErrInvalidNimEmailOrPassword = response.NewError(http.StatusBadRequest, "invalid nim, email or password")
	ErrInitializeAuthRepository  = response.NewError(http.StatusInternalServerError, "failed to initialize auth repository")
)
