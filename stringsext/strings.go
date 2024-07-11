package stringsext

import (
	"unicode"
	"unicode/utf8"

	"math/rand"
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
	l := unicode.ToLower(r)
	if r == l {
		return s
	}
	return string(l) + s[size:]
}

// Rand returns a random string of length l from a given set of runes src
func Rand(src []rune, l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = src[rand.Intn(len(src))]
	}
	return string(b)
}
