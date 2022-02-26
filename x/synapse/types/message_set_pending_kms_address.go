package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetPendingKmsAddress = "set_pending_kms_address"

var _ sdk.Msg = &MsgSetPendingKmsAddress{}

func NewMsgSetPendingKmsAddress(creator string, pendingKmsAddress string) *MsgSetPendingKmsAddress {
	return &MsgSetPendingKmsAddress{
		Creator:           creator,
		PendingKmsAddress: pendingKmsAddress,
	}
}

func (msg *MsgSetPendingKmsAddress) Route() string {
	return RouterKey
}

func (msg *MsgSetPendingKmsAddress) Type() string {
	return TypeMsgSetPendingKmsAddress
}

func (msg *MsgSetPendingKmsAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPendingKmsAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPendingKmsAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
