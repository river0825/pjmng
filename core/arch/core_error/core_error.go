package core_error

import (
	"fmt"

	"github.com/go-errors/errors"
)

type ErrorCode string

type CoreError struct {
	Err  error
	Code ErrorCode
}

func NewCoreError(code ErrorCode, err error) *CoreError {
	return &CoreError{
		Err:  err,
		Code: code,
	}
}

func NewInternalError(msg interface{}) error {
	return errors.New(msg)
}

func NewEntityNotFoundError(entityName string, id interface{}) error {
	return NewCoreError(ErrorCodeEntityNotFound, fmt.Errorf("entity %s with id %v not found", entityName, id))
}

func (e CoreError) Error() string {
	return e.Err.Error()
}

func NewValidateErrorOrNil(err error) error {
	if err == nil {
		return nil
	}

	return NewCoreError(ErrorCodeValidationError, err)
}

func NewInvalidRequest(err string) error {
	return NewCoreError(ErrorCodeInvalidRequest, errors.New(err))
}
