package nominations

import "errors"

var (
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrUniqueViolation     = errors.New("unique constraint violation")
)
