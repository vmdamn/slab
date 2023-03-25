package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateSlab = "create_slab"

var _ sdk.Msg = &MsgCreateSlab{}

func NewMsgCreateSlab(creator string, originatorSocialId string, directedTowards string, assertion string, uriOriginator string) *MsgCreateSlab {
	return &MsgCreateSlab{
		Creator:            creator,
		OriginatorSocialId: originatorSocialId,
		DirectedTowards:    directedTowards,
		Assertion:          assertion,
		UriOriginator:      uriOriginator,
	}
}

func (msg *MsgCreateSlab) Route() string {
	return RouterKey
}

func (msg *MsgCreateSlab) Type() string {
	return TypeMsgCreateSlab
}

func (msg *MsgCreateSlab) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSlab) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSlab) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
