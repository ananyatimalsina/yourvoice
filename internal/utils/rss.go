package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

func GenerateSignature(digest string, key rsa.PrivateKey) (string, error) {
	signature, err := rsa.SignPSS(rand.Reader, &key, crypto.SHA256, []byte(digest), nil)
	if err != nil {
		return "", err
	}
	return string(signature), nil
}

func VerifySignature(digest string, signed string, key rsa.PublicKey) bool {
	return rsa.VerifyPSS(&key, crypto.SHA256, []byte(digest), []byte(signed), nil) == nil
}
