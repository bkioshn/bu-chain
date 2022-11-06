package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExchangeToken = "exchange_token"

var _ sdk.Msg = &MsgExchangeToken{}

func NewMsgExchangeToken(creator string, receiver string, denom string, amount string, exchangeToken string) *MsgExchangeToken {
	return &MsgExchangeToken{
		Creator:       creator,
		Receiver:      receiver,
		Denom:         denom,
		Amount:        amount,
		ExchangeToken: exchangeToken,
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

	amount, err := strconv.ParseUint(msg.Amount, 0, 64)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "amount is not uint", err)
	}
	if amount == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom amount is 0")
	}

	return nil
}
