package types

import (
	sdkmath "cosmossdk.io/math"
	bbn "github.com/Lorenzo-Protocol/lorenzo/v3/types"
)

func NewBTCHeaderInfo(header *bbn.BTCHeaderBytes, headerHash *bbn.BTCHeaderHashBytes, height uint64, work *sdkmath.Uint) *BTCHeaderInfo {
	return &BTCHeaderInfo{
		Header: header,
		Hash:   headerHash,
		Height: height,
		Work:   work,
	}
}

func (m *BTCHeaderInfo) HasParent(parent *BTCHeaderInfo) bool {
	return m.Header.HasParent(parent.Header)
}

func (m *BTCHeaderInfo) Eq(other *BTCHeaderInfo) bool {
	return m.Hash.Eq(other.Hash)
}
