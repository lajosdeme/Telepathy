package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgUnfollowUser struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	UserId  string         `json:"userId" yaml:"userId"`
}

func NewMsgUnfollowUser(creator sdk.AccAddress, userId string) MsgUnfollowUser {
	return MsgUnfollowUser{
		Creator: creator,
		UserId:  userId,
	}
}

func (msg MsgUnfollowUser) Route() string {
	return RouterKey
}

func (msg MsgUnfollowUser) Type() string {
	return "UnfollowUser"
}

func (msg MsgUnfollowUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgUnfollowUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgUnfollowUser) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
