package types_test

import (
	"testing"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/agent/types"
	"github.com/stretchr/testify/suite"
)

type GenesisTestSuite struct {
	suite.Suite
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) TestValidateGenesis() {
	testCases := []struct {
		name     string
		genState *types.GenesisState
		expPass  bool
	}{
		{
			"default",
			types.DefaultGenesisState(),
			true,
		},
		{
			"valid New genesis",
			types.NewGenesisState(
				types.Params{
					AllowList: []string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
					},
				},
				nil,
			),
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.genState.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}
