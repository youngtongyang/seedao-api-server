package ordertoken

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Generate() (plain string, hash string, err error) {
	buf := make([]byte, 32)
	if _, err = rand.Read(buf); err != nil {
		return "", "", fmt.Errorf("generate order token: %w", err)
	}
	plain = "ot_" + base64.RawURLEncoding.EncodeToString(buf)
	hash = Hash(plain)
	return plain, hash, nil
}

func Hash(plain string) string {
	sum := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(sum[:])
}

func Verify(plain, storedHash string) bool {
	if plain == "" || storedHash == "" {
		return false
	}
	return Hash(plain) == storedHash
}
