package util

import (
	"strings"
)

func Crlf2lf(s string) string {
	res := strings.Replace(s, "\r\n", "\n", -1)
	return res
}
