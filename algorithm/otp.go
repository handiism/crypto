package algorithm

import "errors"

// OTP contain a key for One Time Pad algorithm
type OTP struct {
	key string
}

// NewOTP construct One Time Pad algorithm
func NewOTP(key string) *OTP { return &OTP{key} }

// Encrypt encrypts text using a one-time pad
func (otp *OTP) Encrypt(text string) (string, error) {
	if len(otp.key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keyChars := []rune(otp.key)
	return mapAlpha(text, func(i, char int) int {
		return char + alphaIndex(keyChars[i])
	}), nil
}

// Decrypt decrypts text using a one-time pad
func (otp *OTP) Decrypt(text string) (string, error) {
	if len(otp.key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keyChars := []rune(otp.key)
	return mapAlpha(text, func(i, char int) int {
		return char - alphaIndex(keyChars[i])
	}), nil
}
