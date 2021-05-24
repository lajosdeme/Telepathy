package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgDislikeThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgDislikeThought) (*sdk.Result, error) {
	return k.DislikeThought(ctx, msg.ThoughtId, msg.Creator)
}
