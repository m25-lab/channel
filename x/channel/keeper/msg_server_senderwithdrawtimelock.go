package keeper

import (
	"context"
	"errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/m25-lab/channel/x/channel/types"
)

func (k msgServer) Senderwithdrawtimelock(goCtx context.Context, msg *types.MsgSenderwithdrawtimelock) (*types.MsgSenderwithdrawtimelockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	val, found := k.Keeper.GetFwdcommit(ctx, msg.TransferIndex)
	if !found {
		return nil, errors.New("time lock is not existing")
	}

	if val.Sender != msg.To {
		return nil, errors.New("not matching address!")
	}

	timelocksender, err := strconv.ParseUint(val.TimelockSender, 10, 64)
	if err != nil {
		return nil, err
	}
	if timelocksender > uint64(ctx.BlockHeight()) {
		return nil, errors.New("wait until valid blokcheight")
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, sdk.Coins{*val.CoinTransfer})
	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveFwdcommit(ctx, msg.TransferIndex)

	return &types.MsgSenderwithdrawtimelockResponse{}, nil
}
