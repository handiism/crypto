package algorithm

import (
	"fmt"

	"github.com/handiism/crypto/mod"
)

// Super is a key for a Super Encryption
type Super struct {
	A, B int
	Text string
}

// NewSuper creates an Super. For a one-to-one mapping, a must be
// invertable, as in gcd(a, 26) == 1.
func NewSuper(a, b int, text string) (*Super, error) {
	if _, ok := mod.Inverse(a, 26); !ok {
		return nil, fmt.Errorf("no inverse exists for a=%d", a)
	}

	b = mod.Mod(b, 26)
	return &Super{a, b, text}, nil
}

// Encipher enciphers string using Super cipher according to key.
func (key *Super) Encipher(text string) (string, error) {
	caesar := NewCaesar(key.A)
	affine, err := NewAffine(key.A, key.B)
	if err != nil {
		return "", err
	}
	otp := NewOTP(key.Text)

	chiperCaesar := caesar.Encipher(text)
	chiperAffine := affine.Encipher(chiperCaesar)
	chiperOTP, err := otp.Encrypt(chiperAffine)
	if err != nil {
		return "", err
	}

	return chiperOTP, nil
}

// Decipher deciphers string using Super cipher according to key.
func (key *Super) Decipher(chiper string) (string, error) {
	caesar := NewCaesar(key.A)
	affine, err := NewAffine(key.A, key.B)
	if err != nil {
		return "", err
	}
	otp := NewOTP(key.Text)

	plainOTP, err := otp.Decrypt(chiper)
	if err != nil {
		return "", err
	}
	plainAffine := affine.Decipher(plainOTP)
	plainCaesar := caesar.Decipher(plainAffine)

	return plainCaesar, nil
}
