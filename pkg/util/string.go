package util

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
