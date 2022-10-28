package keeper

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"bu-chain/x/exchange/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k msgServer) CreateExchangeRate(goCtx context.Context, msg *types.MsgCreateExchangeRate) (*types.MsgCreateExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check whether the tokens is a pair or not
	tokens := strings.Split(msg.Index, "-")
	if len(tokens) != 2 {
		return nil, types.ErrTokenShouldBePair
	}

	var key string
	rate, err := strconv.ParseFloat(msg.Rate, 64)
	if err != nil {
		return nil, err
	}

	if strings.Compare(tokens[0], tokens[1]) == 1 {
		key = tokens[1] + "-" + tokens[0]
		rate = 1 / rate

	} else {
		key = tokens[0] + "-" + tokens[1]
	}

	// Check if the value already exists
	_, isFound := k.GetExchangeRate(
		ctx,
		key, // msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var exchangeRate = types.ExchangeRate{
		Creator: msg.Creator,
		Index:   key,
		Rate:    fmt.Sprintf("%.3f", rate),
	}

	k.SetExchangeRate(
		ctx,
		exchangeRate,
	)
	return &types.MsgCreateExchangeRateResponse{}, nil
}

func (k msgServer) UpdateExchangeRate(goCtx context.Context, msg *types.MsgUpdateExchangeRate) (*types.MsgUpdateExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetExchangeRate(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var exchangeRate = types.ExchangeRate{
		Creator: msg.Creator,
		Index:   msg.Index,
		Rate:    msg.Rate,
	}

	k.SetExchangeRate(ctx, exchangeRate)

	return &types.MsgUpdateExchangeRateResponse{}, nil
}

func (k msgServer) DeleteExchangeRate(goCtx context.Context, msg *types.MsgDeleteExchangeRate) (*types.MsgDeleteExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetExchangeRate(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveExchangeRate(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteExchangeRateResponse{}, nil
}
