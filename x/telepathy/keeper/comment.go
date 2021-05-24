package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
	"github.com/lajosdeme/telepathy/x/telepathy/utils"
)

// GetCommentCount get the total number of comment
func (k Keeper) GetCommentCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.CommentCountPrefix)
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

// SetCommentCount set the total number of comment
func (k Keeper) SetCommentCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.CommentCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateComment creates a comment
func (k Keeper) CreateComment(ctx sdk.Context, msg types.MsgCreateComment) (*sdk.Result, error) {
	// Create the comment
	count := k.GetCommentCount(ctx)
	var comment = types.Comment{
		Creator:   msg.Creator,
		ID:        strconv.FormatInt(count, 10),
		Message:   msg.Message,
		ThoughtId: msg.ThoughtId,
	}

	thought, err := k.GetThought(ctx, comment.ThoughtId)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Can't add comment to thought.")
	}

	store := ctx.KVStore(k.storeKey)

	//set comment in store
	key := []byte(types.CommentPrefix + comment.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(comment)
	store.Set(key, value)

	// Update comment count
	k.SetCommentCount(ctx, count+1)

	thought.Comments = append(thought.Comments, comment.ID)
	//set thought in store
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(thought)
	storeKey := []byte(types.ThoughtPrefix + thought.ID)
	store.Set(storeKey, bz)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// GetComment returns the comment information
func (k Keeper) GetComment(ctx sdk.Context, key string) (types.Comment, error) {
	store := ctx.KVStore(k.storeKey)
	var comment types.Comment
	byteKey := []byte(types.CommentPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &comment)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

// SetComment sets a comment
func (k Keeper) SetComment(ctx sdk.Context, comment types.Comment) {
	commentKey := comment.ID
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(comment)
	key := []byte(types.CommentPrefix + commentKey)
	store.Set(key, bz)
}

// DeleteComment deletes a comment
func (k Keeper) DeleteComment(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)

	var comment types.Comment
	byteKey := []byte(types.CommentPrefix + key)
	k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &comment)

	thought, _ := k.GetThought(ctx, comment.ThoughtId)
	for i, id := range thought.Comments {
		if id == comment.ID {
			thought.Comments = utils.Remove(thought.Comments, i)
		}
	}
	//set thought in store
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(thought)
	storeKey := []byte(types.ThoughtPrefix + thought.ID)
	store.Set(storeKey, bz)

	store.Delete([]byte(types.CommentPrefix + key))
}

//
// Functions used by querier
//

func listComment(ctx sdk.Context, k Keeper) ([]byte, error) {
	var commentList []types.Comment
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.CommentPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var comment types.Comment
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &comment)
		commentList = append(commentList, comment)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, commentList)
	return res, nil
}

func getComment(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	comment, err := k.GetComment(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, comment)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//List comment for thought id
func listCommentForThought(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	key := path[0]

	var commentList []types.Comment
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.CommentPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var comment types.Comment
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &comment)
		if comment.ThoughtId == key {
			commentList = append(commentList, comment)
		}
	}
	res := codec.MustMarshalJSONIndent(k.cdc, commentList)
	return res, nil
}

// Get creator of the item
func (k Keeper) GetCommentOwner(ctx sdk.Context, key string) sdk.AccAddress {
	comment, err := k.GetComment(ctx, key)
	if err != nil {
		return nil
	}
	return comment.Creator
}

// Check if the key exists in the store
func (k Keeper) CommentExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.CommentPrefix + key))
}

//Like a comment
func (k Keeper) LikeComment(ctx sdk.Context, key string, creator sdk.AccAddress) (*sdk.Result, error) {
	//get store & comment
	store := ctx.KVStore(k.storeKey)
	comment, err := k.GetComment(ctx, key)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrThoughtNotFound, "Cant get comment with provided key.")
	}
	if utils.Contains(comment.Likes, creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Address already likes post.")
	}

	comment.Likes = append(comment.Likes, creator)

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(comment)
	storeKey := []byte(types.CommentPrefix + comment.ID)
	//set in store
	store.Set(storeKey, bz)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

//Dislike a comment
func (k Keeper) DislikeComment(ctx sdk.Context, key string, creator sdk.AccAddress) (*sdk.Result, error) {
	store := ctx.KVStore(k.storeKey)
	comment, err := k.GetComment(ctx, key)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrThoughtNotFound, "Cant get comment with provided key.")
	}

	for i, like := range comment.Likes {
		if like.Equals(creator) {
			comment.Likes = utils.RemoveAddr(comment.Likes, i)
		}
	}

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(comment)
	storeKey := []byte(types.CommentPrefix + comment.ID)
	//set in store
	store.Set(storeKey, bz)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
