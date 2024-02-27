package stringsext

import (
	"unicode"
	"unicode/utf8"
)

func Substring(value string, start int, end int) string {
	return value[start:end]
}

// RemoveFirstChar removes the first character of a string
func RemoveFirstChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func ToLowerFirstChar(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}
