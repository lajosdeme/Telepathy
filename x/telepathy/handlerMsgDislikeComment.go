package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgDislikeComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgDislikeComment) (*sdk.Result, error) {
	return k.DislikeComment(ctx, msg.ThoughtId, msg.Creator)
}
