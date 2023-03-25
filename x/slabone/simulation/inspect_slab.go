package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"slabone/x/slabone/keeper"
	"slabone/x/slabone/types"
)

func SimulateMsgInspectSlab(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgInspectSlab{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the InspectSlab simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "InspectSlab simulation not implemented"), nil, nil
	}
}
