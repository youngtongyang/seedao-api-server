package snsutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	assert.Equal(t, "alice_test.seedao", Normalize("Alice_Test.seedao"))
	assert.Equal(t, "young1.seedao", Normalize("  young1.seedao  "))
}

func TestValidateFormat(t *testing.T) {
	assert.NoError(t, ValidateFormat("alice_test.seedao"))
	assert.Error(t, ValidateFormat("alice_test"))
	assert.Error(t, ValidateFormat("ab.seedao"))
	assert.Error(t, ValidateFormat("invalid-name.seedao"))
}

func TestRegistryLookupNames(t *testing.T) {
	assert.Equal(t, []string{"young.seedao"}, RegistryLookupNames("young.seedao"))
}
