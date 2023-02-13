package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/m25-lab/channel/x/channel/types"
)

func (k msgServer) Receiverwithdraw(goCtx context.Context, msg *types.MsgReceiverwithdraw) (*types.MsgReceiverwithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	val, found := k.Keeper.GetFwdcommit(ctx, msg.TransferIndex)
	if !found {
		return nil, errors.New("commitment is not existing")
	}

	if val.Receiver != msg.To {
		return nil, fmt.Errorf("not matching receiver address! expected:", val.Receiver)
		//return nil, errors.New("not matching receiver address!")
	}

	hash := sha256.Sum256([]byte(msg.Secret))
	if val.Creator == "receiver" {
		// Receiver created this Forward Contract =>
		// the withdrawer could provide either the destination secret or the sender secret
		if val.HashcodeDest != base64.StdEncoding.EncodeToString(hash[:]) &&
			val.HashcodeHtlc != base64.StdEncoding.EncodeToString(hash[:]) {
			return nil, errors.New("Wrong hash !")
		}
	} else {
		// sender created this Forward Contract =>
		// the withdrawer has to provide the destination secret
		if val.HashcodeDest != base64.StdEncoding.EncodeToString(hash[:]) {
			return nil, errors.New("Wrong hash !")
		}
	}

	timelockreceiver, err := strconv.ParseUint(val.TimelockReceiver, 10, 64)
	if err != nil {
		return nil, err
	}
	if timelockreceiver > uint64(ctx.BlockHeight()) {
		return nil, errors.New("wait until valid blokcheight")
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

	return &types.MsgReceiverwithdrawResponse{}, nil
}
