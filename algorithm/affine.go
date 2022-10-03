package algorithm

import (
	"fmt"

	"github.com/handiism/crypto/mod"
)

// Affine is a key for an Affine cipher
type Affine struct {
	A, B, AInv int
}

// NewAffine creates an Affine. For a one-to-one mapping, a must be
// invertable, as in gcd(a, 26) == 1.
func NewAffine(a, b int) (*Affine, error) {
	aInv, ok := mod.Inverse(a, 26)
	if !ok {
		return nil, fmt.Errorf("no inverse exists for a=%d", a)
	}
	b = mod.Mod(b, 26)
	return &Affine{a, b, aInv}, nil
}

// Encipher enciphers string using Affine cipher according to key.
func (key *Affine) Encipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return key.A*char + key.B
	})
}

// Decipher deciphers string using Affine cipher according to key.
func (key *Affine) Decipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return key.AInv * (char - key.B)
	})
}

