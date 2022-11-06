package cli

import (
	"strconv"

	"bu-chain/x/exchange/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdExchangeToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-token [receiver] [denom] [amount] [exchangeToken]",
		Short: "Broadcast message exchange-token",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argDenom := args[1]
			argAmount := args[2]
			argExchangeToken := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgExchangeToken(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argDenom,
				argAmount,
				argExchangeToken,
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

func CmdCreateExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-exchange-rate [pair] [rate] [multiplier]",
		Short: "Create a new exchange-rate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			argPair := args[0]

			// Get value arguments
			argRate, err := strconv.ParseUint(args[1], 0, 64)
			if err != nil {
				return err
			}

			argMul, err := strconv.ParseUint(args[2], 0, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateExchangeRate(
				clientCtx.GetFromAddress().String(),
				argPair,
				argRate,
				argMul,
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

func CmdUpdateExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-exchange-rate [pair] [rate] [multiplier]",
		Short: "Update a exchange-rate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			argPair := args[0]

			// Get value arguments
			argRate, err := strconv.ParseUint(args[1], 0, 64)
			if err != nil {
				return err
			}

			argMul, err := strconv.ParseUint(args[2], 0, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateExchangeRate(
				clientCtx.GetFromAddress().String(),
				argPair,
				argRate,
				argMul,
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

func CmdDeleteExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-exchange-rate [pair]",
		Short: "Delete a exchange-rate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPair := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteExchangeRate(
				clientCtx.GetFromAddress().String(),
				argPair,
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
