package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUser{}

type MsgCreateUser struct {
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	Username string         `json:"username" yaml:"username"`
	Bio      string         `json:"bio" yaml:"bio"`
}

func NewMsgCreateUser(creator sdk.AccAddress, username string, bio string) MsgCreateUser {
	return MsgCreateUser{
		Creator:  creator,
		Username: username,
		Bio:      bio,
	}
}

func (msg MsgCreateUser) Route() string {
	return RouterKey
}

func (msg MsgCreateUser) Type() string {
	return "CreateUser"
}

func (msg MsgCreateUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateUser) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
