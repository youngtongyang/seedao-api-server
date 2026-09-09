package phone

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

const (
	// TestPhonePrefix marks preview/QA numbers that share the production DB.
	// Stats and other production counts must exclude these accounts.
	TestPhonePrefix = "+8619900000"
)

var e164Pattern = regexp.MustCompile(`^\+[1-9]\d{6,14}$`)

func IsTestPhone(phone string) bool {
	return strings.HasPrefix(phone, TestPhonePrefix)
}

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
