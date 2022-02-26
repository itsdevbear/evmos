package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/evmos/x/synapse/types"
)

func (suite *KeeperTestSuite) TestBridgeIn() {
	testCases := []struct {
		name      string
		coins     sdk.Coins
		destAddrs []string
		destEnvs  []string
		malleate  func()
		expPass   bool
	}{
		{
			"single bridge into cosmos - ok",
			sdk.NewCoins(sdk.NewCoin("luna", sdk.NewIntFromUint64(100000))),
			[]string{"evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu"},
			[]string{"cosmos"},
			func() {},
			true,
		},
		{
			"multi bridge into cosmos- ok",
			sdk.NewCoins(sdk.NewCoin("luna", sdk.NewIntFromUint64(500000)), sdk.NewCoin("bera", sdk.NewIntFromUint64(400000)), sdk.NewCoin("atom", sdk.NewIntFromUint64(300000))),
			[]string{"evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu", "evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu", "evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu"},
			[]string{"cosmos", "cosmos", "cosmos"},
			func() {},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.mintFeeCollector = true
			suite.SetupTest()
			tc.malleate()
			suite.Commit()
			data := make([]*types.InBridgeData, len(tc.coins))

			for i := range tc.coins {
				data[i] = &types.InBridgeData{
					Coin:     &tc.coins[i],
					DestAddr: tc.destAddrs[i],
					DestEnv:  tc.destEnvs[i],
				}
			}

			msg := types.NewMsgBridgeIn(
				suite.address.String(),
				data,
			)

			err := suite.app.SynapseKeeper.HandleBridgeIn(suite.ctx, msg)
			suite.Commit()
			if tc.expPass {
				suite.Require().NoError(err, tc.name)
				for i := range tc.coins {
					accAddress, _ := sdk.AccAddressFromBech32(tc.destAddrs[i])
					suite.Require().Equal(tc.coins[i].Amount.Uint64(), suite.app.BankKeeper.GetBalance(suite.ctx, accAddress, tc.coins[i].Denom).Amount.Uint64())
				}
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
}
