package snsutil

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	ReasonInvalidFormat   = "invalid_format"
	ReasonTakenInApp      = "taken_in_app"
	ReasonTakenInRegistry = "taken_in_registry"
	ReasonReservedWord    = "reserved_word"

	SeedaoSuffix = ".seedao"
)

var seedaoNamePattern = regexp.MustCompile(`^[a-z0-9_]{3,20}\.seedao$`)

// Normalize lowercases and trims; input must already include .seedao (see ValidateFormat).
func Normalize(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

// ValidateFormat requires canonical form: {base}.seedao
func ValidateFormat(name string) error {
	if !seedaoNamePattern.MatchString(Normalize(name)) {
		return fmt.Errorf(ReasonInvalidFormat)
	}
	return nil
}

// RegistryLookupNames returns names used to match sns_chain_registry rows.
func RegistryLookupNames(name string) []string {
	return []string{Normalize(name)}
}
