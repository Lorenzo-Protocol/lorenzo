package keeper_test

func (suite *KeeperTestSuite) TestEndBlocker() {
	suite.SetupTest()
	oldParams := suite.keeper.GetParams(suite.ctx)
	suite.Require().Empty(oldParams.Beacon)
	suite.keeper.EndBlocker(suite.ctx)

	newParams := suite.keeper.GetParams(suite.ctx)
	suite.Require().NotEmpty(newParams.Beacon)
}
