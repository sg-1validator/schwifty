package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateNft = "create_nft"
	TypeMsgUpdateNft = "update_nft"
	TypeMsgDeleteNft = "delete_nft"
)

var _ sdk.Msg = &MsgCreateNft{}

func NewMsgCreateNft(creator string, collectionId uint64, owner string, uri string, uriHash string, data string) *MsgCreateNft {
	return &MsgCreateNft{
		Creator:      creator,
		CollectionId: collectionId,
		Owner:        owner,
		Uri:          uri,
		UriHash:      uriHash,
		Data:         data,
	}
}

func (msg *MsgCreateNft) Route() string {
	return RouterKey
}

func (msg *MsgCreateNft) Type() string {
	return TypeMsgCreateNft
}

func (msg *MsgCreateNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNft{}

func NewMsgUpdateNft(creator string, id uint64, collectionId uint64, owner string, uri string, uriHash string, data string) *MsgUpdateNft {
	return &MsgUpdateNft{
		Id:           id,
		Creator:      creator,
		CollectionId: collectionId,
		Owner:        owner,
		Uri:          uri,
		UriHash:      uriHash,
		Data:         data,
	}
}

func (msg *MsgUpdateNft) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNft) Type() string {
	return TypeMsgUpdateNft
}

func (msg *MsgUpdateNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteNft{}

func NewMsgDeleteNft(creator string, id uint64) *MsgDeleteNft {
	return &MsgDeleteNft{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteNft) Route() string {
	return RouterKey
}

func (msg *MsgDeleteNft) Type() string {
	return TypeMsgDeleteNft
}

func (msg *MsgDeleteNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
