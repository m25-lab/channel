package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/m25-lab/channel/testutil/keeper"
	"github.com/m25-lab/channel/testutil/nullify"
	"github.com/m25-lab/channel/x/channel/keeper"
	"github.com/m25-lab/channel/x/channel/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNChannel(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Channel {
	items := make([]types.Channel, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetChannel(ctx, items[i])
	}
	return items
}

func TestChannelGet(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNChannel(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChannel(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestChannelRemove(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNChannel(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChannel(ctx,
			item.Index,
		)
		_, found := keeper.GetChannel(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestChannelGetAll(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNChannel(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllChannel(ctx)),
	)
}
