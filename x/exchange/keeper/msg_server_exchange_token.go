package keeper

import (
	"context"
	"math"
	"strconv"
	"strings"

	"bu-chain/x/exchange/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ExchangeToken(goCtx context.Context, msg *types.MsgExchangeToken) (*types.MsgExchangeTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	// Check err
	owner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	amount := sdk.NewCoins(msg.Denom)

	exchangePair := msg.Denom.Denom + "-" + msg.ExchangeDenom
	rate, isFound := k.GetExchangeRate(ctx, exchangePair)
	exRate, err := strconv.ParseFloat(rate.Rate, 64)

	if strings.Compare(msg.Denom.Denom, msg.ExchangeDenom) == 1 {
		exchangePair = msg.ExchangeDenom + "-" + msg.Denom.Denom
		rate, isFound = k.GetExchangeRate(ctx, exchangePair)
		if !isFound {
			return nil, types.ErrTokenPairNotFound
		}
		exRate, err = strconv.ParseFloat(rate.Rate, 64)
		if err != nil {
			return nil, err
		}

		exRate = 1 / exRate
	}
	if !isFound {
		return nil, types.ErrTokenPairNotFound
	}
	if err != nil {
		return nil, err
	}

	exRate64 := uint64(exRate * math.Pow10(3))

	if !isFound {
		return nil, types.ErrTokenPairNotFound
	}

	// Receiver get amount
	err = k.bankKeeper.SendCoins(ctx, owner, receiver, amount)
	if err != nil {
		return nil, err
	}

	// Owner get amount
	exchangeAmount := sdk.Coins{sdk.NewInt64Coin(msg.ExchangeDenom, msg.Denom.Amount.Int64()*int64(exRate64)/int64(math.Pow10(3)))}
	err = k.bankKeeper.SendCoins(ctx, receiver, owner, exchangeAmount)
	if err != nil {
		return nil, err
	}

	return &types.MsgExchangeTokenResponse{}, nil
}
