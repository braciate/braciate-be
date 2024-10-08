package assets

import "github.com/braciate/braciate-be/internal/pkg/response"

var (
	ErrForeignKeyViolation = response.NewError(400, "foreign key violation")
	ErrUniqueViolation     = response.NewError(400, "unique constraint violation")
)
