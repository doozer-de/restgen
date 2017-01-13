package registry

import (
	"unicode"
	"unicode/utf8"
)

func isIn(s string, in []string) bool {
	for _, x := range in {
		if s == x {
			return true
		}
	}
	return false
}

func firstCharacterToLower(s string) string {
	if s == "" {
		return ""
	}

	r, n := utf8.DecodeRuneInString(s)

	return string(unicode.ToLower(r)) + s[n:]
}

func firstCharacterToUpper(s string) string {
	if s == "" {
		return ""
	}

	r, n := utf8.DecodeRuneInString(s)

	return string(unicode.ToUpper(r)) + s[n:]
}
