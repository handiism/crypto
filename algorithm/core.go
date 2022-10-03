package algorithm

import (
	"regexp"

	"github.com/handiism/crypto/mod"
)

func mapAlpha(text string, f func(i, char int) int) string {
	runes := []rune(text)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = rune(mod.Mod(f(i, int(char-'A')), 26)) + 'A'
		} else if char >= 'a' && char <= 'z' {
			runes[i] = rune(mod.Mod(f(i, int(char-'a')), 26)) + 'a'
		}
	}
	return string(runes)
}

// RemovePunctuation removes any non alphabetic chars from a string.
func RemovePunctuation(text string) string {
	return replacePattern(text, "[^A-Za-z]", "")
}

// replacePattern replaces any matches to a Regular Expression in a string.
func replacePattern(text, pattern, replace string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(text, replace)
}
