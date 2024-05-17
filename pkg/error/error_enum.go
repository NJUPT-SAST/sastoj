package error

import (
	"errors"
	"fmt"
)

type LocalError struct {
	Code     int
	GRPCCode int
	Message  string
	Err      error
}

func (e LocalError) Error() string {
	return fmt.Sprintf("ErrCode: %d, ErrMsg: %s, Err: %v", e.Code, e.Message, e.Err)
}

var (
	UnknownError = LocalError{Code: 100, Message: "未知错误"}
	ServerError  = LocalError{Code: 101, Message: "服务器内部错误"}
)

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
