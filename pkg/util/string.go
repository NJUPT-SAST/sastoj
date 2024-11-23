package util

import (
	"strconv"
	"strings"
)

func BytesMatchIgnoringSpacesAndNewlines(a, b []byte) bool {
	skipWhitespace := func(c byte) bool {
		return c == ' ' || c == '\n' || c == '\t' || c == '\r'
	}
	i, j := 0, 0
	lenA, lenB := len(a), len(b)
	for i < lenA && j < lenB {
		for i < lenA && skipWhitespace(a[i]) {
			i++
		}
		for j < lenB && skipWhitespace(b[j]) {
			j++
		}
		if i < lenA && j < lenB {
			if a[i] != b[j] {
				return false
			}
			i++
			j++
		}
	}
	for i < lenA && skipWhitespace(a[i]) {
		i++
	}
	for j < lenB && skipWhitespace(b[j]) {
		j++
	}
	return i == lenA && j == lenB
}

func ParseInt64(s string) int64 {
	parseInt, _ := strconv.ParseInt(s, 10, 64)
	return parseInt
}

func GetCaseIndex(s string) int {
	res := 0
	for _, b := range []byte(s) {
		if b >= '0' && b <= '9' {
			res = res*10 + int(b-'0')
		} else {
			break
		}
	}
	return res
}

func StringCompareIgnoreLineEndSpaceAndTextEndEnter(a, b string) bool {
	a = strings.TrimRight(a, " \n\t\r")
	b = strings.TrimRight(b, " \n\t\r")
	a = StringRemoveLineEndSpace(a)
	b = StringRemoveLineEndSpace(b)
	return a == b
}

func StringRemoveLineEndSpace(a string) string {
	// Convert the string to a byte slice to avoid frequent string allocations
	b := []byte(a)
	n := len(b)
	result := make([]byte, 0, n) // Preallocate space to improve performance

	for i := 0; i < n; {
		// Find the position of the next newline character
		j := i
		for j < n && b[j] != '\n' {
			j++
		}

		// Remove trailing spaces before the newline
		end := j
		for end > i && b[end-1] == ' ' {
			end--
		}

		// Append the cleaned line to the result
		result = append(result, b[i:end]...)

		// Move to the start of the next line
		i = j + 1
	}

	return string(result)
}

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
