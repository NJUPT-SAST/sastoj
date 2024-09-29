package error

import (
	"errors"

	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// FromError : Convert kratos error to Local-error , mainly add customized error code
func FromError(err error) *LocalError {
	var kratosError *kerr.Error
	kratosError = kerr.FromError(err)
	if kratosError.Status.Reason == "VALIDATOR" {
		return &LocalError{
			Code:    InvalidParamError.Code,
			Message: kratosError.Message,
			Reason:  InvalidParamError.Reason,
			Err:     kratosError,
		}
	}
	if _, ok := errorMap[kratosError.Reason]; ok {
		return &LocalError{
			Code:    errorMap[kratosError.Reason],
			Message: kratosError.Message,
			Reason:  kratosError.Reason,
			Err:     err,
		}
	}
	log.Debug("error not found in errorMap")
	return &LocalError{
		Code:    ServerError.Code,
		Message: ServerError.Message,
		Reason:  ServerError.Reason,
		Err:     err,
	}

}

func (e *LocalError) Wrap(err error) LocalError {
	e.Err = err
	return *e
}

// Is : determine whether the error is equal
func (e *LocalError) Is(err error) bool {
	var localError LocalError
	if errors.As(err, &localError) {
		return localError.Code == e.Code
	}
	return false
}

// HandleError :this func is used to handle error
// When a function has multiple errors,
// instead of using if-else to determine them one by one,
// use this function to get the errors.
func HandleError(err error) LocalError {
	var localError LocalError
	if errors.As(err, &localError) {
		return err.(LocalError)
	}
	// if not exist, return default error
	return ServerError.Wrap(err)
}
