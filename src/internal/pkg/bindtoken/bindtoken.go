package bindtoken

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

const bindCodeChars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
const bindCodeLength = 6

func GenerateBindToken() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("generate bind token: %w", err)
	}
	return "fbt_" + base64.RawURLEncoding.EncodeToString(buf), nil
}

func GenerateBindCode() (string, error) {
	max := big.NewInt(int64(len(bindCodeChars)))
	code := make([]byte, bindCodeLength)
	for i := range code {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", fmt.Errorf("generate bind code: %w", err)
		}
		code[i] = bindCodeChars[n.Int64()]
	}
	return string(code), nil
}
