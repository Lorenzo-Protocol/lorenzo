package keeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMintToAddr(t *testing.T) {
	assert := assert.New(t)
	_, err := validateMintToAddress([]byte("0xc1912fee45d61c87cc5ea59dae31190fffff232d"))
	assert.NoError(err, "should be valid")
	_, err = validateMintToAddress([]byte("c1912fee45d61c87cc5ea59dae31190fffff232d"))
	assert.NoError(err, "should be valid")
	_, err = validateMintToAddress([]byte("0XC1912FEE45D61C87CC5EA59DAE31190FFFFF232D"))
	assert.NoError(err, "should be valid")
	_, err = validateMintToAddress([]byte("0xC1912fEE45d61C87Cc5EA59DaE31190FFFFf232d"))
	assert.Error(err, "should be invalid")
	_, err = validateMintToAddress([]byte("0x7c52e508C07558C287d5A453475954f6a547eC41"))
	assert.NoError(err, "should be valid")
	_, err = validateMintToAddress([]byte("0xc1912fEE45d61C87Cc5EA59DaE31190FFFFf232d"))
	assert.NoError(err, "should be valid")
}
