package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/slabone module sentinel errors

var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrWrongSlabState = sdkerrors.Register(ModuleName, 5100, "It is an inappropriate current Slab state.") //Only Created Slabs can be Vetted and Vetted Slabs can be Revoked.
	ErrVetterIsOriginator = sdkerrors.Register(ModuleName, 5200, "Is an Originator. Vetter can't be Originator")
	ErrNotDirectedTowards = sdkerrors.Register(ModuleName, 5300, "Is where it should be directed towards. Not directed towards the Vetter")
	ErrRevokerShouldBeVetter = sdkerrors.Register(ModuleName, 5400, "Should be Revoker i.e. Vetter Should Be Revoker")
)
