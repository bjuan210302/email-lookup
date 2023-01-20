package model

import (
	"errors"
	"fmt"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func NewRequestError(value int, errMessage string) *RequestError {
	err := errors.New(errMessage)
	return &RequestError{
		StatusCode: value,
		Err:        err,
	}
}

func (ve *RequestError) Error() string {
	return fmt.Sprint(ve.Err.Error())
}
