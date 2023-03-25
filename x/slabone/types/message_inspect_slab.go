package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInspectSlab = "inspect_slab"

var _ sdk.Msg = &MsgInspectSlab{}

func NewMsgInspectSlab(creator string, vetterSocialId string, vettingNote string, uriVetter string, id uint64) *MsgInspectSlab {
	return &MsgInspectSlab{
		Creator:        creator,
		VetterSocialId: vetterSocialId,
		VettingNote:    vettingNote,
		UriVetter:      uriVetter,
		Id:             id,
	}
}

func (msg *MsgInspectSlab) Route() string {
	return RouterKey
}

func (msg *MsgInspectSlab) Type() string {
	return TypeMsgInspectSlab
}

func (msg *MsgInspectSlab) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInspectSlab) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInspectSlab) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
