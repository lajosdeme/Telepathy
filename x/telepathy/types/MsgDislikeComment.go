package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgDislikeComment struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	ThoughtId string         `json:"commentId" yaml:"commentId"`
}

func NewMsgDislikeComment(creator sdk.AccAddress, thoughtId string) MsgDislikeComment {
	return MsgDislikeComment{
		Creator:   creator,
		ThoughtId: thoughtId,
	}
}

func (msg MsgDislikeComment) Route() string {
	return RouterKey
}

func (msg MsgDislikeComment) Type() string {
	return "DislikeComment"
}

func (msg MsgDislikeComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDislikeComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDislikeComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
