package blog_test

import (
	"testing"

	keepertest "github.com/cosmonaut/planet/testutil/keeper"
	"github.com/cosmonaut/planet/x/blog"
	"github.com/cosmonaut/planet/x/blog/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogKeeper(t)
	blog.InitGenesis(ctx, *k, genesisState)
	got := blog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Equal(t, genesisState.PortId, got.PortId)
	// this line is used by starport scaffolding # genesis/test/assert
}
