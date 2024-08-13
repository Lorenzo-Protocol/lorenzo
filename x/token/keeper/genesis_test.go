package keeper_test

import "github.com/Lorenzo-Protocol/lorenzo/v3/x/token/types"

func (suite *KeeperTestSuite) TestInitGenesis() {
	testcases := []struct {
		name string
		gs   types.GenesisState
	}{
		{
			name: "empty genesis state",
			gs:   types.GenesisState{},
		},
		{
			name: "genesis state with token pairs",
			gs: types.GenesisState{
				TokenPairs: []types.TokenPair{
					{
						ContractAddress: "0x1D1530e3A3719BE0BEe1abba5016Cf2e236f3277",
						Denom:           coinBaseDenom,
						Source:          types.OWNER_MODULE,
						Enabled:         true,
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.Require().NotPanics(func() {
				suite.app.TokenKeeper.InitGenesis(suite.ctx, &tc.gs)
			})
			params := suite.app.TokenKeeper.GetParams(suite.ctx)
			tokenPairs := suite.app.TokenKeeper.GetTokenPairs(suite.ctx)
			suite.Require().Equal(tc.gs.Params, params)

			if len(tokenPairs) > 0 {
				suite.Require().Equal(tc.gs.TokenPairs, tokenPairs)
			} else {
				suite.Require().Len(tc.gs.TokenPairs, 0)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestExportGenesis() {
	testcases := []struct {
		name string
		gs   types.GenesisState
	}{
		{
			name: "empty genesis",
			gs:   types.GenesisState{},
		},
		{
			name: "default genesis",
			gs:   *types.DefaultGenesisState(),
		},
		{
			name: "genesis state with token pairs",
			gs: types.GenesisState{
				TokenPairs: []types.TokenPair{
					{
						ContractAddress: "0x1D1530e3A3719BE0BEe1abba5016Cf2e236f3277",
						Denom:           coinBaseDenom,
						Source:          types.OWNER_MODULE,
						Enabled:         true,
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.app.TokenKeeper.InitGenesis(suite.ctx, &tc.gs)

			suite.Require().NotPanics(func() {
				gs := suite.app.TokenKeeper.ExportGenesis(suite.ctx)

				params := suite.app.TokenKeeper.GetParams(suite.ctx)
				suite.Require().Equal(gs.Params, params)

				tokenPairs := suite.app.TokenKeeper.GetTokenPairs(suite.ctx)
				if len(tokenPairs) > 0 {
					suite.Require().Equal(gs.TokenPairs, tokenPairs)
				} else {
					suite.Require().Len(gs.TokenPairs, 0)
				}
			})
		})
	}
}
