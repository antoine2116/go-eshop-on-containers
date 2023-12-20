package custom_errors

import "net/http"

type internalServerError struct {
	CustomError
}

type InternalServerError interface {
	CustomError
}

func NewInternalServerError(err error) error {
	return &internalServerError{
		NewCustomError(http.StatusInternalServerError, "", err),
	}
}
