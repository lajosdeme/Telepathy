package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

/* -------------------------------------------------------------------------- */
/* --------------------------- Create user message -------------------------- */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/* ---------------------------- Set user message ---------------------------- */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/* --------------------------- Delete user message -------------------------- */
/* -------------------------------------------------------------------------- */

type MsgDeleteUser struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteUser(id string, creator sdk.AccAddress) MsgDeleteUser {
	return MsgDeleteUser{
		ID:      id,
		Creator: creator,
	}
}

func (msg MsgDeleteUser) Route() string {
	return RouterKey
}

func (msg MsgDeleteUser) Type() string {
	return "DeleteUser"
}

func (msg MsgDeleteUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteUser) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

/* -------------------------------------------------------------------------- */
/* --------------------------- Follow user message -------------------------- */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/* -------------------------- Unfollow user message ------------------------- */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/* --------------------------- Set avatar message --------------------------- */
/* -------------------------------------------------------------------------- */

type MsgSetAvatar struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Avatar  string         `json:"avatar" yaml:"avatar"`
}

func NewMsgSetAvatar(creator sdk.AccAddress, id string, avatar string) MsgSetAvatar {
	return MsgSetAvatar{
		Creator: creator,
		ID:      id,
		Avatar:  avatar,
	}
}

func (msg MsgSetAvatar) Route() string {
	return RouterKey
}

func (msg MsgSetAvatar) Type() string {
	return "SetAvatar"
}

func (msg MsgSetAvatar) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetAvatar) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetAvatar) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Avatar == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "avatar hash can't be empty")
	}
	return nil
}
