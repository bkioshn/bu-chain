package keeper

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"bu-chain/x/exchange/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateExchangeRate(goCtx context.Context, msg *types.MsgCreateExchangeRate) (*types.MsgCreateExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check whether the tokens is a pair or not
	tokens := strings.Split(msg.Index, "-")
	if len(tokens) != 2 {
		return nil, errors.New("Input tokens need to be in pair")
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
