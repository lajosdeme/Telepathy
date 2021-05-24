package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrMsgTooLong        = sdkerrors.Register(ModuleName, 1, "Message must be under 250 characters.")
	ErrThoughtNotFound   = sdkerrors.Register(ModuleName, 2, "Thought with given key doesn't exist.")
	ErrUserAlreadyExists = sdkerrors.Register(ModuleName, 3, "User for creator address already exists.")
)
