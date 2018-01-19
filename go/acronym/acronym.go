package acronym

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const testVersion = 3

// Abbreviate returns an acronym for a string.
func Abbreviate(s string) string {
	res := ""
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(s, f)
	for _, word := range words {
		rune, _ := utf8.DecodeRuneInString(word)
		upperRune := unicode.ToUpper(rune)
		res += string(upperRune)
	}
	return res
}
