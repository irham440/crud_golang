package utils

import "errors"

type Errors struct {
	Message string
}

func (e *Errors) Error() string {
	return e.Message
}

func NewErrors(message string) error {
	return &Errors{Message: message}
}

func NewError(message string) error {
	return errors.New(message)
}
