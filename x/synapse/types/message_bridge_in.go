package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBridgeIn = "bridge_in"

var _ sdk.Msg = &MsgBridgeIn{}

func NewMsgBridgeIn(creator string, data []*InBridgeData) *MsgBridgeIn {
	return &MsgBridgeIn{
		Creator: creator,
		Data:    data,
	}
}

func (msg *MsgBridgeIn) Route() string {
	return RouterKey
}

func (msg *MsgBridgeIn) Type() string {
	return TypeMsgBridgeIn
}

func (msg *MsgBridgeIn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBridgeIn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBridgeIn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
