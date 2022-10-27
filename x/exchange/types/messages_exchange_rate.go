package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateExchangeRate = "create_exchange_rate"
	TypeMsgUpdateExchangeRate = "update_exchange_rate"
	TypeMsgDeleteExchangeRate = "delete_exchange_rate"
)

var _ sdk.Msg = &MsgCreateExchangeRate{}

func NewMsgCreateExchangeRate(
	creator string,
	index string,
	rate string,

) *MsgCreateExchangeRate {
	return &MsgCreateExchangeRate{
		Creator: creator,
		Index:   index,
		Rate:    rate,
	}
}

func (msg *MsgCreateExchangeRate) Route() string {
	return RouterKey
}

func (msg *MsgCreateExchangeRate) Type() string {
	return TypeMsgCreateExchangeRate
}

func (msg *MsgCreateExchangeRate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateExchangeRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateExchangeRate{}

func NewMsgUpdateExchangeRate(
	creator string,
	index string,
	rate string,

) *MsgUpdateExchangeRate {
	return &MsgUpdateExchangeRate{
		Creator: creator,
		Index:   index,
		Rate:    rate,
	}
}

func (msg *MsgUpdateExchangeRate) Route() string {
	return RouterKey
}

func (msg *MsgUpdateExchangeRate) Type() string {
	return TypeMsgUpdateExchangeRate
}

func (msg *MsgUpdateExchangeRate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateExchangeRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteExchangeRate{}

func NewMsgDeleteExchangeRate(
	creator string,
	index string,

) *MsgDeleteExchangeRate {
	return &MsgDeleteExchangeRate{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteExchangeRate) Route() string {
	return RouterKey
}

func (msg *MsgDeleteExchangeRate) Type() string {
	return TypeMsgDeleteExchangeRate
}

func (msg *MsgDeleteExchangeRate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteExchangeRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
