package algorithm

import (
	"github.com/handiism/crypto/mod"
)

// mapAlpha returns the shifted text for defined callback function
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

// alphaIndex returns the number 0-25 corresponding to the letter
func alphaIndex(char rune) int {
	if char >= 'A' && char <= 'Z' {
		return int(char - 'A')
	} else if char >= 'a' && char <= 'z' {
		return int(char - 'a')
	}
	return -1
}
