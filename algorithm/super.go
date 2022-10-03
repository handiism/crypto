package algorithm

import (
	"fmt"

	"github.com/handiism/crypto/mod"
)

// Super is a key for a Super Encryption
type Super struct {
	A, B int
}

// NewSuper creates an Super. For a one-to-one mapping, a must be
// invertable, as in gcd(a, 26) == 1.
func NewSuper(a, b int) (*Super, error) {
	if _, ok := mod.Inverse(a, 26); !ok {
		return nil, fmt.Errorf("no inverse exists for a=%d", a)
	}

	b = mod.Mod(b, 26)
	return &Super{a, b}, nil
}

// Encipher enciphers string using Super cipher according to key.
func (key *Super) Encipher(text string) (string, error) {
	caesar := NewCaesar(key.A)
	affine, _ := NewAffine(key.A, key.B)
	railfence, err := NewRailfence(key.A)
	if err != nil {
		return "", err
	}
	chiperCaesar := caesar.Encipher(text)
	chiperAffine := affine.Encipher(chiperCaesar)
	chiperRailfence, err := railfence.Encipher(chiperAffine)
	if err != nil {
		return "", err
	}

	return chiperRailfence, nil
}

// Decipher deciphers string using Super cipher according to key.
func (key *Super) Decipher(chiper string) (string, error) {
	caesar := NewCaesar(key.A)
	affine, _ := NewAffine(key.A, key.B)
	railfence, err := NewRailfence(key.A)
	if err != nil {
		return "", err
	}
	plainRailfence, err := railfence.Decipher(chiper)
	if err != nil {
		return "", err
	}
	plainAffine := affine.Decipher(plainRailfence)
	plainCaesar := caesar.Decipher(plainAffine)

	return plainCaesar, nil
}
