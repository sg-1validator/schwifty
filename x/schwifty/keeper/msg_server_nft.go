package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

func (k msgServer) CreateNft(goCtx context.Context, msg *types.MsgCreateNft) (*types.MsgCreateNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var nft = types.Nft{
		Creator:      msg.Creator,
		CollectionId: msg.CollectionId,
		Owner:        msg.Owner,
		Uri:          msg.Uri,
		UriHash:      msg.UriHash,
		Data:         msg.Data,
	}
	
	collection, found := k.GetCollection(ctx, nft.CollectionId)
	if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "key %d doesn't exist")
    }

	if msg.Creator != collection.Owner {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Not the owner of the collection, cannot create new NFTs here")
    }

	id := k.AppendNft(
		ctx,
		nft,
	)

	return &types.MsgCreateNftResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateNft(goCtx context.Context, msg *types.MsgUpdateNft) (*types.MsgUpdateNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var nft = types.Nft{
		Creator:      msg.Creator,
		Id:           msg.Id,
		CollectionId: msg.CollectionId,
		Owner:        msg.Owner,
		Uri:          msg.Uri,
		UriHash:      msg.UriHash,
		Data:         msg.Data,
	}

	collection, found := k.GetCollection(ctx, nft.CollectionId)
	if !found {
        return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
    }

	if msg.Creator != collection.Owner {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Not the owner of the collection, cannot create new NFTs here")
    }

	// Checks that the element exists
	val, found := k.GetNft(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetNft(ctx, nft)

	return &types.MsgUpdateNftResponse{}, nil
}

func (k msgServer) DeleteNft(goCtx context.Context, msg *types.MsgDeleteNft) (*types.MsgDeleteNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetNft(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	collection, found := k.GetCollection(ctx, val.CollectionId)
	if !found {
        return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "collection key %d doesn't exist", val.CollectionId)
    }

	if msg.Creator != collection.Owner {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Not the owner of the collection, cannot create new NFTs here")
    }

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveNft(ctx, msg.Id)

	return &types.MsgDeleteNftResponse{}, nil
}
