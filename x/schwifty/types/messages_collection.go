package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCollection = "create_collection"
	TypeMsgUpdateCollection = "update_collection"
	TypeMsgDeleteCollection = "delete_collection"
)

var _ sdk.Msg = &MsgCreateCollection{}

func NewMsgCreateCollection(creator string, owner string, name string, description string, ticker string, uri string, uriHash string, data string) *MsgCreateCollection {
	return &MsgCreateCollection{
		Creator:     creator,
		Owner:       owner,
		Name:        name,
		Description: description,
		Ticker:      ticker,
		Uri:         uri,
		UriHash:     uriHash,
		Data:        data,
	}
}

func (msg *MsgCreateCollection) Route() string {
	return RouterKey
}

func (msg *MsgCreateCollection) Type() string {
	return TypeMsgCreateCollection
}

func (msg *MsgCreateCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCollection{}

func NewMsgUpdateCollection(creator string, id uint64, owner string, name string, description string, ticker string, uri string, uriHash string, data string) *MsgUpdateCollection {
	return &MsgUpdateCollection{
		Id:          id,
		Creator:     creator,
		Owner:       owner,
		Name:        name,
		Description: description,
		Ticker:      ticker,
		Uri:         uri,
		UriHash:     uriHash,
		Data:        data,
	}
}

func (msg *MsgUpdateCollection) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCollection) Type() string {
	return TypeMsgUpdateCollection
}

func (msg *MsgUpdateCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCollection{}

func NewMsgDeleteCollection(creator string, id uint64) *MsgDeleteCollection {
	return &MsgDeleteCollection{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteCollection) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCollection) Type() string {
	return TypeMsgDeleteCollection
}

func (msg *MsgDeleteCollection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCollection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
