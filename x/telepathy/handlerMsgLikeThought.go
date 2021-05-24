package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgLikeThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgLikeThought) (*sdk.Result, error) {
	return k.LikeThought(ctx, msg.ThoughtId, msg.Creator)
}
