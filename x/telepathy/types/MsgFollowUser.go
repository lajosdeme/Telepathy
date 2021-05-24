package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgFollowUser struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	UserId  string         `json:"userId" yaml:"userId"`
}

func NewMsgFollowUser(creator sdk.AccAddress, userId string) MsgFollowUser {
	return MsgFollowUser{
		Creator: creator,
		UserId:  userId,
	}
}

func (msg MsgFollowUser) Route() string {
	return RouterKey
}

func (msg MsgFollowUser) Type() string {
	return "FollowUser"
}

func (msg MsgFollowUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgFollowUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgFollowUser) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
