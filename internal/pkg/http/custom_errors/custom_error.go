package custom_errors

import (
	"errors"
	"fmt"
)

type customError struct {
	statusCode int
	message    string
	err        error
}

type CustomError interface {
	error
	GetStatus() int
	GetMessage() string
}

func (e *customError) Error() string {
	return fmt.Sprintf("message: %s, error: %v", e.message, e.err.Error())
}

func (e *customError) GetStatus() int {
	return e.statusCode
}

func (e *customError) GetMessage() string {
	return e.message
}

func NewCustomError(statusCode int, message string, err error) CustomError {
	return &customError{
		statusCode: statusCode,
		message:    message,
		err:        err,
	}
}

func GetCustomError(err error) CustomError {
	var customError CustomError
	if errors.As(err, &customError) {
		return customError
	}
	return nil
}
