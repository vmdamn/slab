package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokeSlab = "revoke_slab"

var _ sdk.Msg = &MsgRevokeSlab{}

func NewMsgRevokeSlab(creator string, revokerSocialId string, revokingNote string, uriRevoker string, id uint64) *MsgRevokeSlab {
	return &MsgRevokeSlab{
		Creator:         creator,
		RevokerSocialId: revokerSocialId,
		RevokingNote:    revokingNote,
		UriRevoker:      uriRevoker,
		Id:              id,
	}
}

func (msg *MsgRevokeSlab) Route() string {
	return RouterKey
}

func (msg *MsgRevokeSlab) Type() string {
	return TypeMsgRevokeSlab
}

func (msg *MsgRevokeSlab) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeSlab) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeSlab) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
