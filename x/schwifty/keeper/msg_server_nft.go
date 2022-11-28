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

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveNft(ctx, msg.Id)

	return &types.MsgDeleteNftResponse{}, nil
}
