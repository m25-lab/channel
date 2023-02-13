package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/m25-lab/channel/x/channel/types"
)

func (k msgServer) WithdrawHashlock(goCtx context.Context, msg *types.MsgWithdrawHashlock) (*types.MsgWithdrawHashlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	val, found := k.Keeper.GetCommitment(ctx, msg.Index)
	if !found {
		return nil, errors.New("commitment is not existing")
	}

	if val.ToHashlockAddr != msg.To {
		return nil, fmt.Errorf("not matching receiver address! expected:", val.ToHashlockAddr)
		//return nil, errors.New("not matching receiver address!")
	}

	hash := sha256.Sum256([]byte(msg.Secret))
	if val.Hashcode != base64.StdEncoding.EncodeToString(hash[:]) {
		return nil, errors.New("Wrong hash !")
	}

	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, sdk.Coins{*val.CoinToHtlc})
	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveCommitment(ctx, msg.Index)

	return &types.MsgWithdrawHashlockResponse{}, nil
}
