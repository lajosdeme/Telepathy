package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgLikeComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgLikeComment) (*sdk.Result, error) {
	return k.LikeComment(ctx, msg.CommentId, msg.Creator)
}
