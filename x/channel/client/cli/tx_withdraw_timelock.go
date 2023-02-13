package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/m25-lab/channel/x/channel/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWithdrawTimelock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-timelock [to] [index]",
		Short: "Broadcast message withdraw-timelock",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTo := args[0]
			argIndex := args[1]

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawTimelock(
				clientCtx.GetFromAddress().String(),
				argTo,
				argIndex,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
