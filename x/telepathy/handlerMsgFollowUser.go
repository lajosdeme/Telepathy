package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgFollowUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgFollowUser) (*sdk.Result, error) {
	return k.FollowUser(ctx, msg.UserId, msg.Creator)
}
