package telepathy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgCreateUser:
			return handleMsgCreateUser(ctx, k, msg)
		case types.MsgSetUser:
			return handleMsgSetUser(ctx, k, msg)
		case types.MsgDeleteUser:
			return handleMsgDeleteUser(ctx, k, msg)
		case types.MsgCreateComment:
			return handleMsgCreateComment(ctx, k, msg)
		case types.MsgSetComment:
			return handleMsgSetComment(ctx, k, msg)
		case types.MsgDeleteComment:
			return handleMsgDeleteComment(ctx, k, msg)
		case types.MsgCreateThought:
			return handleMsgCreateThought(ctx, k, msg)
		case types.MsgSetThought:
			return handleMsgSetThought(ctx, k, msg)
		case types.MsgDeleteThought:
			return handleMsgDeleteThought(ctx, k, msg)
		case types.MsgLikeThought:
			return handleMsgLikeThought(ctx, k, msg)
		case types.MsgLikeComment:
			return handleMsgLikeComment(ctx, k, msg)
		case types.MsgDislikeThought:
			return handleMsgDislikeThought(ctx, k, msg)
		case types.MsgDislikeComment:
			return handleMsgDislikeComment(ctx, k, msg)
		case types.MsgFollowUser:
			return handleMsgFollowUser(ctx, k, msg)
		case types.MsgUnfollowUser:
			return handleMsgUnfollowUser(ctx, k, msg)
		case types.MsgSetAvatar:
			return handleMsgSetAvatar(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
