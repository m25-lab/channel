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

func (k msgServer) Senderwithdrawhashlock(goCtx context.Context, msg *types.MsgSenderwithdrawhashlock) (*types.MsgSenderwithdrawhashlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	val, found := k.Keeper.GetFwdcommit(ctx, msg.TransferIndex)
	if !found {
		return nil, errors.New("commitment is not existing")
	}

	if val.Sender != msg.To {
		return nil, fmt.Errorf("not matching receiver address! expected:", val.Sender)
		//return nil, errors.New("not matching receiver address!")
	}

	hash := sha256.Sum256([]byte(msg.Secret))
	if val.HashcodeHtlc != base64.StdEncoding.EncodeToString(hash[:]) {
		return nil, errors.New("Wrong hash !")
	}

	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, sdk.Coins{*val.CoinTransfer})
	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveFwdcommit(ctx, msg.TransferIndex)

	return &types.MsgSenderwithdrawhashlockResponse{}, nil
}
