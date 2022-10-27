package cli

import (
	"strconv"

	"bu-chain/x/exchange/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdExchangeToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-token [receiver] [denom] [exchangeDenom]",
		Short: "Broadcast message exchange-token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argDenom, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			argExchangeDenom := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgExchangeToken(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argDenom,
				argExchangeDenom,
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
