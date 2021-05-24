package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgUnfollowUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgUnfollowUser) (*sdk.Result, error) {
	return k.UnfollowUser(ctx, msg.UserId, msg.Creator)
}
