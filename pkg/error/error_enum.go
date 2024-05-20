package error

import (
	"fmt"
)

type LocalError struct {
	Code    int
	Reason  string
	Message string
	Err     error
}

func (e LocalError) Error() string {
	return fmt.Sprintf("ErrCode: %d, ErrMsg: %s, Err: %v", e.Code, e.Message, e.Err)
}

var (
	// define common and universal errors

	UnknownError      = LocalError{Code: 1000, Reason: "UNKNOWN_ERROR", Message: "未知错误"}
	ServerError       = LocalError{Code: 1001, Reason: "SERVER_ERROR", Message: "服务器内部错误"}
	InvalidParamError = LocalError{Code: 1002, Reason: "INVALID_ARGUMENT", Message: "参数错误"}

	// define contest modulo errors

	ContestNotFoundError = LocalError{Code: 2001, Reason: "CONTEST_NOT_FOUND", Message: "比赛不存在"}
	ContestInvalidError  = LocalError{Code: 2002, Reason: "CONTEST_INVALID", Message: "比赛无效"}
	ContestExistError    = LocalError{Code: 2003, Reason: "CONTEST_EXIST", Message: "比赛已存在"}

	GroupNotFoundError = LocalError{Code: 3001, Reason: "GROUP_NOT_FOUND", Message: "群组不存在"}
	GroupInvalidError  = LocalError{Code: 3002, Reason: "GROUP_INVALID", Message: "群组无效"}
	GroupExistError    = LocalError{Code: 3003, Reason: "GROUP_EXIST", Message: "群组已存在"}

	UserNotFoundError = LocalError{Code: 4001, Reason: "USER_NOT_FOUND", Message: "用户不存在"}
	UserExistError    = LocalError{Code: 4002, Reason: "USER_EXIST", Message: "用户已存在"}
)

var errorMap = map[string]int{
	"UNKNOWN_ERROR":     1000,
	"SERVER_ERROR":      1001,
	"INVALID_ARGUMENT":  1002,
	"CONTEST_NOT_FOUND": 2001,
	"CONTEST_INVALID":   2002,
	"CONTEST_EXIST":     2003,
	"GROUP_NOT_FOUND":   3001,
	"GROUP_INVALID":     3002,
	"GROUP_EXIST":       3003,
	"USER_NOT_FOUND":    4001,
	"USER_EXIST":        4002,
}
