package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/m25-lab/channel/x/channel/keeper"
	"github.com/m25-lab/channel/x/channel/types"
)

func SimulateMsgOpenChannel(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOpenChannel{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the OpenChannel simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "OpenChannel simulation not implemented"), nil, nil
	}
}
