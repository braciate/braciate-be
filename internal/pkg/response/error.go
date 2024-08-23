package response

type Error struct {
	Code int
	Err  string
}

func (e *Error) Error() string {
	return e.Err
}

func NewError(code int, err string) error {
	return &Error{code, err}
}
