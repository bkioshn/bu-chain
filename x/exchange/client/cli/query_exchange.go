package cli

import (
	"context"

	"bu-chain/x/exchange/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-exchange-rate",
		Short: "list all exchange-rate",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllExchangeRateRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ExchangeRateAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-exchange-rate [index]",
		Short: "shows an exchange-rate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetExchangeRateRequest{
				Index: argIndex,
			}

			res, err := queryClient.ExchangeRate(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdExchangeAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-amount [denom] [amount] [exchange-token]",
		Short: "Query exchange-amount",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqAmount := args[1]
			reqExchangeToken := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryExchangeAmountRequest{

				Denom:         reqDenom,
				Amount:        reqAmount,
				ExchangeToken: reqExchangeToken,
			}

			res, err := queryClient.ExchangeAmount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdExchangePairs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-pairs",
		Short: "list all exchange pairs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryExchangePairsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ExchangePairs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd

}
