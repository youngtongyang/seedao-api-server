package ordertoken_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/pkg/ordertoken"
)

func TestGenerateAndVerify(t *testing.T) {
	plain, hash, err := ordertoken.Generate()
	require.NoError(t, err)
	assert.True(t, ordertoken.Verify(plain, hash))
	assert.False(t, ordertoken.Verify("ot_wrong", hash))
}
