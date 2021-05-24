package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetUser{}

type MsgSetUser struct {
	ID       string         `json:"id" yaml:"id"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	Username string         `json:"username" yaml:"username"`
	Bio      string         `json:"bio" yaml:"bio"`
}

func NewMsgSetUser(creator sdk.AccAddress, id string, username string, bio string) MsgSetUser {
	return MsgSetUser{
		ID:       id,
		Creator:  creator,
		Username: username,
		Bio:      bio,
	}
}

func (msg MsgSetUser) Route() string {
	return RouterKey
}

func (msg MsgSetUser) Type() string {
	return "SetUser"
}

func (msg MsgSetUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetUser) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
