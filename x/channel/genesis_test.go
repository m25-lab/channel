package channel_test

import (
	"testing"

	keepertest "github.com/m25-lab/channel/testutil/keeper"
	"github.com/m25-lab/channel/testutil/nullify"
	"github.com/m25-lab/channel/x/channel"
	"github.com/m25-lab/channel/x/channel/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CommitmentList: []types.Commitment{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ChannelList: []types.Channel{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		FwdcommitList: []types.Fwdcommit{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChannelKeeper(t)
	channel.InitGenesis(ctx, *k, genesisState)
	got := channel.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CommitmentList, got.CommitmentList)
	require.ElementsMatch(t, genesisState.ChannelList, got.ChannelList)
	require.ElementsMatch(t, genesisState.FwdcommitList, got.FwdcommitList)
	// this line is used by starport scaffolding # genesis/test/assert
}
