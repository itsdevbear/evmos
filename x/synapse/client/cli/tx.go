package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/evmos/x/synapse/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
	_                                     = strconv.Itoa(0)
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdBridgeIn())
	cmd.AddCommand(CmdBridgeOut())
	cmd.AddCommand(CmdSetPendingKmsAddress())
	cmd.AddCommand(CmdSetKmsAddress())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdBridgeIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge-in [coins] [dest-addrs] [dest-envs]",
		Short: "Broadcast message BridgeIn",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCoins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}
			argDestAddrs := strings.Split(args[1], listSeparator)
			argDestEnvs := strings.Split(args[2], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			data := make([]*types.InBridgeData, len(argCoins))

			for i := range argCoins {
				data[i] = &types.InBridgeData{
					Coin:     &argCoins[i],
					DestAddr: argDestAddrs[i],
					DestEnv:  argDestEnvs[i],
				}
			}

			msg := types.NewMsgBridgeIn(
				clientCtx.GetFromAddress().String(),
				data,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdBridgeOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge-out [coins] [dest-addrs] [dest-chains]",
		Short: "Broadcast message BridgeOut",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCoins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}
			argDestAddrs := strings.Split(args[1], listSeparator)
			argDestChains := strings.Split(args[2], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			data := make([]*types.OutBridgeData, len(argCoins))

			for i := range argCoins {
				data[i] = &types.OutBridgeData{
					Coin:      &argCoins[i],
					DestAddr:  argDestAddrs[i],
					DestChain: argDestChains[i],
				}
			}

			msg := types.NewMsgBridgeOut(
				clientCtx.GetFromAddress().String(),
				data,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
