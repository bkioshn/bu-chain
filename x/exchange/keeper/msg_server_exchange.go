package keeper

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	appparams "bu-chain/app/params"
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

	uMultiplier := 1e6
	uDenom := "u" + strings.ToLower(msg.Denom)
	denomAmount, err := strconv.ParseFloat(msg.Amount, 64)

	if err != nil {
		return nil, err
	}

	uDenomAmount := denomAmount * uMultiplier
	denomCoin := sdk.NewCoins(sdk.NewCoin(uDenom, sdk.NewInt(int64(uDenomAmount))))

	exchangePair := msg.Denom + "-" + msg.ExchangeToken
	var tokenReceivedAmount uint64
	if msg.Denom == appparams.DisplayDenom {
		exchangePair = msg.ExchangeToken + "-" + msg.Denom
		rate, isFound := k.GetExchangeRate(ctx, exchangePair)
		if !isFound {
			return nil, types.ErrTokenPairNotFound
		}
		tokenReceivedAmount = uint64(uDenomAmount) * rate.Multiplier / rate.Rate
	} else {
		rate, isFound := k.GetExchangeRate(ctx, exchangePair)
		if !isFound {
			return nil, types.ErrTokenPairNotFound
		}
		tokenReceivedAmount = uint64(denomAmount * float64(rate.Rate))
	}

	uExchangeToken := "u" + strings.ToLower(msg.ExchangeToken)
	if k.bankKeeper.GetBalance(ctx, receiver, uExchangeToken).Amount.LT(sdk.NewInt(int64(tokenReceivedAmount))) {
		return nil, errors.New("Receiver doen't have enough amount to exchange")
	}

	// Receiver get amount
	err = k.bankKeeper.SendCoins(ctx, owner, receiver, denomCoin)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Owner send %f", uDenomAmount)

	// Owner get amount
	exchangeCoin := sdk.Coins{sdk.NewCoin(uExchangeToken, sdk.NewInt(int64(tokenReceivedAmount)))}
	err = k.bankKeeper.SendCoins(ctx, receiver, owner, exchangeCoin)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Receiver send %d", tokenReceivedAmount)

	return &types.MsgExchangeTokenResponse{}, nil
}

func (k msgServer) CreateExchangeRate(goCtx context.Context, msg *types.MsgCreateExchangeRate) (*types.MsgCreateExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	key := msg.Index + "-" + appparams.DisplayDenom
	// Check if the value already exists
	_, isFound := k.GetExchangeRate(
		ctx,
		key, // msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "pair already set")
	}

	var exchangeRate = types.ExchangeRate{
		Creator:    msg.Creator,
		Index:      key,
		Rate:       msg.Rate,
		Time:       uint64(ctx.BlockTime().Unix()),
		Multiplier: msg.Multiplier,
	}

	k.SetExchangeRate(
		ctx,
		exchangeRate,
	)
	return &types.MsgCreateExchangeRateResponse{}, nil
}

func (k msgServer) UpdateExchangeRate(goCtx context.Context, msg *types.MsgUpdateExchangeRate) (*types.MsgUpdateExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	key := msg.Index + "-" + appparams.DisplayDenom

	// Check if the value exists
	valFound, isFound := k.GetExchangeRate(
		ctx,
		key,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "pair not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var exchangeRate = types.ExchangeRate{
		Creator:    msg.Creator,
		Index:      key,
		Rate:       msg.Rate,
		Time:       uint64(ctx.BlockTime().Unix()),
		Multiplier: msg.Multiplier,
	}

	k.SetExchangeRate(ctx, exchangeRate)

	return &types.MsgUpdateExchangeRateResponse{}, nil
}

func (k msgServer) DeleteExchangeRate(goCtx context.Context, msg *types.MsgDeleteExchangeRate) (*types.MsgDeleteExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	key := msg.Index + "-" + appparams.DisplayDenom
	// Check if the value exists
	valFound, isFound := k.GetExchangeRate(
		ctx,
		key,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "pair not set")
	}

	// Checks if the the msg creator is the same as the current owner
	//TODO - Fix
	if msg.Creator != valFound.Creator && valFound.Creator != "buchain" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveExchangeRate(
		ctx,
		key,
	)

	return &types.MsgDeleteExchangeRateResponse{}, nil
}
