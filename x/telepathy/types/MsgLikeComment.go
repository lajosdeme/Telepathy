package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgLikeComment struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	CommentId string         `json:"commentId" yaml:"CommentId"`
}

func NewMsgLikeComment(creator sdk.AccAddress, commentId string) MsgLikeComment {
	return MsgLikeComment{
		Creator:   creator,
		CommentId: commentId,
	}
}

func (msg MsgLikeComment) Route() string {
	return RouterKey
}

func (msg MsgLikeComment) Type() string {
	return "LikeComment"
}

func (msg MsgLikeComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgLikeComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgLikeComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
