package util

import "strconv"

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
