package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestROT13Encipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var expected = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
	actual := ROT13.Encipher(text)
	assert.Equal(t, expected, actual)
}

func TestROT13Decipher(t *testing.T) {
	var text = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
	var expected = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	actual := ROT13.Decipher(text)
	assert.Equal(t, expected, actual)
}
