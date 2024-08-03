package keeper_test

import "github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"

func (suite *KeeperTestSuite) TestUploadHeaders() {
	headers := types.GetTestHeaders(suite.T())
	_, err := suite.msgServer.UploadHeaders(suite.ctx, &types.MsgUploadHeaders{
		Headers: headers[5:],
		Signer:  testAdmin.String(),
	})
	suite.Require().NoError(err)
}