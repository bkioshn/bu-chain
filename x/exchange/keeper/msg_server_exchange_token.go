package keeper

import (
	"context"
	"errors"

	"bu-chain/x/exchange/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ExchangeToken(goCtx context.Context, msg *types.MsgExchangeToken) (*types.MsgExchangeTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	owner, _ := sdk.AccAddressFromBech32(msg.Creator)
	receiver, _ := sdk.AccAddressFromBech32(msg.Receiver)
	amount := sdk.Coins{msg.Denom}

	exchangeDenom := k.bankKeeper.GetBalance(ctx, receiver, msg.ExchangeDenom)
	if exchangeDenom.Amount.LT(msg.Denom.Amount) {
		return nil, errors.New(msg.Denom.Amount.String() + ", " + exchangeDenom.Amount.String() + " not enough amount")
	}

	err := k.bankKeeper.SendCoins(ctx, owner, receiver, amount)
	if err != nil {
		return nil, err
	}

	exchangeAmount := sdk.Coins{sdk.NewInt64Coin(msg.ExchangeDenom, msg.Denom.Amount.Int64())}
	// Change name
	err2 := k.bankKeeper.SendCoins(ctx, receiver, owner, exchangeAmount)
	if err2 != nil {
		return nil, err2
	}

	return &types.MsgExchangeTokenResponse{}, nil
}
