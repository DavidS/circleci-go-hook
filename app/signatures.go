package app

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func VerifySignature(body []byte, key []byte, signature_hex string) bool {
	signature, err := hex.DecodeString(signature_hex)
	if err != nil {
		panic(err)
	}

	mac := hmac.New(sha256.New, key)
	mac.Write(body)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(signature, expectedMAC)
}
