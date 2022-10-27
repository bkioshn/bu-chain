package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExchangeToken = "exchange_token"

var _ sdk.Msg = &MsgExchangeToken{}

func NewMsgExchangeToken(creator string, receiver string, denom sdk.Coin, exchangeDenom string) *MsgExchangeToken {
	return &MsgExchangeToken{
		Creator:       creator,
		Receiver:      receiver,
		Denom:         denom,
		ExchangeDenom: exchangeDenom,
	}
}

func (msg *MsgExchangeToken) Route() string {
	return RouterKey
}

func (msg *MsgExchangeToken) Type() string {
	return TypeMsgExchangeToken
}

func (msg *MsgExchangeToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExchangeToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExchangeToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, errReceiver := sdk.AccAddressFromBech32(msg.Receiver)
	if errReceiver != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", errReceiver)
	}

	amount, _ := sdk.ParseCoinsNormalized(msg.Denom.String())
	if !amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom is not a valid Coins object")
	}
	if amount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom amount is empty")
	}

	return nil
}
