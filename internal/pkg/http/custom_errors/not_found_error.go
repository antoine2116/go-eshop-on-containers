package custom_errors

import "net/http"

type notFoundError struct {
	CustomError
}

type NotFoundError interface {
	CustomError
}

func NewNotFoundError(message string, err error) error {
	return &notFoundError{
		NewCustomError(http.StatusNotFound, message, err),
	}
}
