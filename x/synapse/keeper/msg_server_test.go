package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	ethermint "github.com/tharsis/ethermint/types"
	"github.com/tharsis/evmos/contracts"
	erc20types "github.com/tharsis/evmos/x/erc20/types"
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
		//{
		//	"bad destEnv - err",
		//	sdk.NewCoins(sdk.NewCoin("luna", sdk.NewIntFromUint64(100000))),
		//	[]string{"evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu"},
		//	[]string{"bad"},
		//	func() {},
		//	false,
		//},
		{
			"single bridge into cosmos - ok",
			sdk.NewCoins(sdk.NewCoin("luna", sdk.NewIntFromUint64(100000))),
			[]string{"evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu"},
			[]string{"cosmos"},
			func() {},
			true,
		},
		{
			"single bridge into evm w/bech32 address - ok",
			sdk.NewCoins(sdk.NewCoin("luna", sdk.NewIntFromUint64(30000))),
			[]string{"evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu"},
			[]string{"evm"},
			func() {},
			true,
		},
		{
			"single bridge into evm w/hex address - ok",
			sdk.NewCoins(sdk.NewCoin("luna", sdk.NewIntFromUint64(100000))),
			[]string{"0xE019b48E08a238EbE9f7107f8327B60803a4bA67"},
			[]string{"evm"},
			func() {},
			true,
		},
		{
			"multi bridge into single address cosmos- ok",
			sdk.NewCoins(sdk.NewCoin("luna", sdk.NewIntFromUint64(500000)), sdk.NewCoin("bera", sdk.NewIntFromUint64(400000)), sdk.NewCoin("atom", sdk.NewIntFromUint64(300000))),
			[]string{"evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu", "evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu", "evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu"},
			[]string{"cosmos", "cosmos", "cosmos"},
			func() {},
			true,
		},
		{
			"multi bridge multi address into cosmos- ok",
			sdk.NewCoins(sdk.NewCoin("luna", sdk.NewIntFromUint64(500000)), sdk.NewCoin("bera", sdk.NewIntFromUint64(400000)), sdk.NewCoin("atom", sdk.NewIntFromUint64(300000))),
			[]string{"evmos1g9rcsa8sr6utp073pdtvqg5xvynv00vs6dcyfp", "evmos1ssfzktgqclfjtq7aey08pg38w2v6sas3qmkhnu", "evmos1uqvmfrsg5guwh60hzplcxfakpqp6fwn86f9pxd"},
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

			metadata := banktypes.Metadata{
				Description: "LUNA",
				Base:        "luna",
				// NOTE: Denom units MUST be increasing
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "luna",
						Exponent: 0,
					},
				},
				Name:    erc20types.CreateDenom("LUNA"),
				Symbol:  "LUNA",
				Display: "LUNA",
			}

			// Required for ERC20 Module invariant
			suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("luna", sdk.NewInt(1))))

			// Register Tokens
			lunaPair, err := suite.app.Erc20Keeper.RegisterCoin(suite.ctx, metadata)
			if !tc.expPass {
				suite.Require().Error(err, tc.name)
				return
			}

			err = suite.app.SynapseKeeper.HandleBridgeIn(suite.ctx, msg)
			if !tc.expPass {
				suite.Require().Error(err, tc.name)
				return
			}

			suite.Commit()

			for i := range tc.coins {
				if tc.destEnvs[i] == "cosmos" {
					if tc.expPass {
						suite.Require().NoError(err, tc.name)
						accAddress, _ := sdk.AccAddressFromBech32(tc.destAddrs[i])
						suite.Require().Equal(tc.coins[i].Amount.Uint64(), suite.app.BankKeeper.GetBalance(suite.ctx, accAddress, tc.coins[i].Denom).Amount.Uint64())
					} else {
						suite.Require().Error(err, tc.name)

					}
				} else if tc.destEnvs[i] == "evm" {
					if tc.expPass {
						suite.Require().NoError(err, tc.name)
						fmt.Println(lunaPair, tc.destAddrs[i])
						sdk.AccAddressFromBech32(tc.destAddrs[i])
						addr := common.HexToAddress(tc.destAddrs[i])
						if !common.IsHexAddress(tc.destAddrs[i]) {
							acc, _ := sdk.AccAddressFromBech32(tc.destAddrs[i])
							addr = suite.app.AccountKeeper.GetAccount(suite.ctx, acc).(ethermint.EthAccountI).EthAddress()
						}
						balance := suite.app.Erc20Keeper.BalanceOf(suite.ctx, contracts.CosmosRelayedERC20Contract.ABI, common.HexToAddress(lunaPair.Erc20Address), addr)
						fmt.Println(balance.Uint64())
						suite.Require().Equal(tc.coins[i].Amount.Uint64(), balance.Uint64())
					} else {
						suite.Require().Error(err, tc.name)

					}
				}
			}
		})
	}
}
