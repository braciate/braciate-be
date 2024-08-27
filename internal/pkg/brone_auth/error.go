package broneAuth

import (
	"github.com/braciate/braciate-be/internal/pkg/response"
	"net/http"
)

var (
	ErrInvalidEmailNimOrPassword    = response.NewError(http.StatusBadRequest, "invalid email, nim or password")
	ErrResponseNotContainsSAML      = response.NewError(http.StatusInternalServerError, "SAML response from brone not contains saml")
	ErrUnexpected                   = response.NewError(http.StatusInternalServerError, "unexpected error")
	ErrMakingRequestToIam           = response.NewError(http.StatusInternalServerError, "failed to make request to iam")
	ErrMakingRequestToBrone         = response.NewError(http.StatusInternalServerError, "failed to make request to brone")
	ErrMissingRequiredFieldResponse = response.NewError(http.StatusExpectationFailed, "missing required attributes in SAML response")
)
