package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgCreateUser{}, "telepathy/CreateUser", nil)
	cdc.RegisterConcrete(MsgSetUser{}, "telepathy/SetUser", nil)
	cdc.RegisterConcrete(MsgDeleteUser{}, "telepathy/DeleteUser", nil)
	cdc.RegisterConcrete(MsgCreateComment{}, "telepathy/CreateComment", nil)
	cdc.RegisterConcrete(MsgSetComment{}, "telepathy/SetComment", nil)
	cdc.RegisterConcrete(MsgDeleteComment{}, "telepathy/DeleteComment", nil)
	cdc.RegisterConcrete(MsgCreateThought{}, "telepathy/CreateThought", nil)
	cdc.RegisterConcrete(MsgSetThought{}, "telepathy/SetThought", nil)
	cdc.RegisterConcrete(MsgDeleteThought{}, "telepathy/DeleteThought", nil)
	cdc.RegisterConcrete(MsgLikeThought{}, "telepathy/LikeTought", nil)
	cdc.RegisterConcrete(MsgLikeComment{}, "telepathy/LikeComment", nil)
	cdc.RegisterConcrete(MsgDislikeThought{}, "telepathy/DislikeThought", nil)
	cdc.RegisterConcrete(MsgDislikeComment{}, "telepathy/DislikeComment", nil)
	cdc.RegisterConcrete(MsgFollowUser{}, "telepathy/FollowUser", nil)
	cdc.RegisterConcrete(MsgUnfollowUser{}, "telepathy/UnfollowUser", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
