package exception

type InternalServerError struct {
	Error error
}

// day 5 week 4
// error handling 1

func NewInternalServerError(err error) InternalServerError {
	return InternalServerError{err}
}
