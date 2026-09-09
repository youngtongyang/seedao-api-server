package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateCode_DevFixedCodeSkipsSend(t *testing.T) {
	s := NewSMSService(nil, nil, nil, true)
	require.NoError(t, s.ValidateCode(context.Background(), "+8613800138000", SMSDevFixedCode))
	require.Error(t, s.ValidateCode(context.Background(), "13800138000", SMSDevFixedCode))
}
