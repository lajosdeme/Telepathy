package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteThought{}

type MsgDeleteThought struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteThought(id string, creator sdk.AccAddress) MsgDeleteThought {
	return MsgDeleteThought{
		ID:      id,
		Creator: creator,
	}
}

func (msg MsgDeleteThought) Route() string {
	return RouterKey
}

func (msg MsgDeleteThought) Type() string {
	return "DeleteThought"
}

func (msg MsgDeleteThought) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteThought) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteThought) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
