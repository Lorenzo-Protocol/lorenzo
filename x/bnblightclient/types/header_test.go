package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyHeaders(t *testing.T) {
	headers := GetTestHeaders(t)
	require.NoError(t, VerifyHeaders(headers))
}