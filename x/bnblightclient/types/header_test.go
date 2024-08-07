package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/testutil"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
)

func TestVerifyHeaders(t *testing.T) {
	headers := testutil.GetTestHeaders(t)
	require.NoError(t, types.VerifyHeaders(headers))
}