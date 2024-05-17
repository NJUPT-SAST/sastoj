package error

import (
	"errors"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// FromError : Convert kratos error to Local-error , mainly add customized error code
func FromError(err error) *LocalError {
	var kerror *kerr.Error
	kerror = kerr.FromError(err)
	if _, ok := errorMap[kerror.Reason]; ok {
		return &LocalError{
			Code:    errorMap[kerror.Reason],
			Message: kerror.Message,
			Reason:  kerror.Reason,
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
	var lerr LocalError
	if errors.As(err, &lerr) {
		return lerr.Code == e.Code
	}
	return false
}

// HandleError :this func is used to handle error
// When a function has multiple errors,
// instead of using if-else to determine them one by one,
// use this function to get the errors.
func HandleError(err error) LocalError {
	var lerr LocalError
	if errors.As(err, &lerr) {
		return err.(LocalError)
	}
	// if not exist, return default error
	return ServerError.Wrap(err)
}
