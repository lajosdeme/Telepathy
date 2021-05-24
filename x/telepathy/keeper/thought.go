package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
	"github.com/lajosdeme/telepathy/x/telepathy/utils"
)

// GetThoughtCount get the total number of thought
func (k Keeper) GetThoughtCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ThoughtCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetThoughtCount set the total number of thought
func (k Keeper) SetThoughtCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ThoughtCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateThought creates a thought
func (k Keeper) CreateThought(ctx sdk.Context, msg types.MsgCreateThought) {
	// Create the thought
	count := k.GetThoughtCount(ctx)
	var thought = types.Thought{
		Creator: msg.Creator,
		ID:      strconv.FormatInt(count, 10),
		Message: msg.Message,
	}

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ThoughtPrefix + thought.ID)

	value := k.cdc.MustMarshalBinaryLengthPrefixed(thought)
	store.Set(key, value)

	// Update thought count
	k.SetThoughtCount(ctx, count+1)
}

// GetThought returns the thought information
func (k Keeper) GetThought(ctx sdk.Context, key string) (types.Thought, error) {
	store := ctx.KVStore(k.storeKey)
	var thought types.Thought
	byteKey := []byte(types.ThoughtPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &thought)
	if err != nil {
		return thought, err
	}
	return thought, nil
}

// SetThought sets a thought
func (k Keeper) SetThought(ctx sdk.Context, thought types.Thought) {
	thoughtKey := thought.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(thought)
	key := []byte(types.ThoughtPrefix + thoughtKey)
	store.Set(key, bz)
}

// DeleteThought deletes a thought
func (k Keeper) DeleteThought(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ThoughtPrefix + key))
}

func (k Keeper) ListThoughtByCreator(ctx sdk.Context, creatorAddr sdk.AccAddress) []types.Thought {
	var thoughtList []types.Thought
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ThoughtPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var thought types.Thought
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &thought)
		if thought.Creator.Equals(creatorAddr) {
			thoughtList = append(thoughtList, thought)
		}
	}
	return thoughtList
}

//
// Functions used by querier
//

func listThought(ctx sdk.Context, k Keeper) ([]byte, error) {
	var thoughtList []types.Thought
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ThoughtPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var thought types.Thought
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &thought)
		thoughtList = append(thoughtList, thought)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, thoughtList)
	return res, nil
}

func getThought(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	thought, err := k.GetThought(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, thought)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//List all thoughts by creator
func listThoughtByCreator(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	creatorAddrStr := path[0]
	creatorAddr, err := sdk.AccAddressFromBech32(creatorAddrStr)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, creatorAddrStr)
	}

	var thoughtList []types.Thought
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ThoughtPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var thought types.Thought
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &thought)
		if thought.Creator.Equals(creatorAddr) {
			thoughtList = append(thoughtList, thought)
		}
	}
	res := codec.MustMarshalJSONIndent(k.cdc, thoughtList)
	return res, nil
}

// Get creator of the item
func (k Keeper) GetThoughtOwner(ctx sdk.Context, key string) sdk.AccAddress {
	thought, err := k.GetThought(ctx, key)
	if err != nil {
		return nil
	}
	return thought.Creator
}

// Check if the key exists in the store
func (k Keeper) ThoughtExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ThoughtPrefix + key))
}

//Like a thought
func (k Keeper) LikeThought(ctx sdk.Context, key string, creator sdk.AccAddress) (*sdk.Result, error) {
	//get store & thought
	store := ctx.KVStore(k.storeKey)
	thought, err := k.GetThought(ctx, key)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrThoughtNotFound, "Cant get thought with provided key.")
	}
	if utils.Contains(thought.Likes, creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Address already likes post.")
	}
	//set appending creator to likes
	thought.Likes = append(thought.Likes, creator)
	//marshal to byte slice

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(thought)
	storeKey := []byte(types.ThoughtPrefix + thought.ID)
	//set in store
	store.Set(storeKey, bz)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

//Dislike a thought
func (k Keeper) DislikeThought(ctx sdk.Context, key string, creator sdk.AccAddress) (*sdk.Result, error) {
	store := ctx.KVStore(k.storeKey)
	thought, err := k.GetThought(ctx, key)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrThoughtNotFound, "Cant get thought with provided key.")
	}

	for i, like := range thought.Likes {
		if like.Equals(creator) {
			thought.Likes = utils.RemoveAddr(thought.Likes, i)
		}
	}

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(thought)
	storeKey := []byte(types.ThoughtPrefix + thought.ID)
	//set in store
	store.Set(storeKey, bz)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
