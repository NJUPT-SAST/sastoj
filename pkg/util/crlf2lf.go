package util

import (
	"strings"
)

func Crlf2lf(s string) string {
	res := strings.Replace(s, "\r\n", "\n", -1)
	return res
}

func RemoveCr(s []byte) []byte {
	var result []byte
	for _, b := range s {
		if b != '\r' {
			result = append(result, b)
		}
	}
	return result
}

func RemoveLf(s []byte) []byte {
	var result []byte
	for _, b := range s {
		if b != '\n' {
			result = append(result, b)
		}
	}
	return result
}
