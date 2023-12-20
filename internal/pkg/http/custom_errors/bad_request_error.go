package custom_errors

import "net/http"

type badRequestError struct {
	CustomError
}

type BadRequestError interface {
	CustomError
}

func NewBadRequestError(message string, err error) error {
	return &badRequestError{
		NewCustomError(http.StatusBadRequest, message, err),
	}
}
