package phone

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

var e164Pattern = regexp.MustCompile(`^\+[1-9]\d{6,14}$`)

func Validate(phone string) error {
	if !e164Pattern.MatchString(phone) {
		return fmt.Errorf("invalid phone format")
	}
	return nil
}

func Hash(phone string) string {
	sum := sha256.Sum256([]byte(phone))
	return hex.EncodeToString(sum[:])
}

func Mask(phone string) string {
	if len(phone) < 8 {
		return phone
	}
	prefix := phone[:len(phone)-8]
	suffix := phone[len(phone)-4:]
	return prefix + "****" + suffix
}

func NormalizeWallet(addr string) string {
	return strings.ToLower(strings.TrimSpace(addr))
}
