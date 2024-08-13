package keeper_test

import (
	"math/big"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/testutil"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
	"github.com/ethereum/go-ethereum/common"
)

func (suite *KeeperTestSuite) TestVerifyReceipt() {
	data := testutil.GetTestProvedReceipts(suite.T())
	err := suite.keeper.VerifyReceipt(suite.ctx, data.Number, data.Receipt, data.Proof)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestParseReceipt() {
	data := testutil.GetTestProvedReceipts(suite.T())
	events, err := suite.keeper.ParseReceipt(suite.ctx, data.Receipt)
	suite.Require().NoError(err)

	expectEvent := types.CrossChainEvent{
		ChainID:            56,
		Contract:           common.HexToAddress("0x9ADb675bc89d9EC5d829709e85562b7c99658D59"),
		Identifier:         0,
		Sender:             common.HexToAddress("0xdFb41DC2173D2Be024e6d64a83fD011d4ae43E01"),
		PlanID:             1,
		BTCcontractAddress: common.HexToAddress("0x49fF00552CA23899ba9f814bCf7eD55bC5cDd9Ce"),
		StakeAmount:        *big.NewInt(2000000000000000000),
		StBTCAmount:        *big.NewInt(2000000000000000000),
	}
	suite.Require().Len(events, 1, "expected 1 event")
	suite.Require().EqualValues(expectEvent, events[0], "event mismatch")
}
