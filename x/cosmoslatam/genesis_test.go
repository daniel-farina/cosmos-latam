package cosmoslatam_test

import (
	"testing"

	keepertest "github.com/daniel-farina/cosmos-latam/testutil/keeper"
	"github.com/daniel-farina/cosmos-latam/x/cosmoslatam"
	"github.com/daniel-farina/cosmos-latam/x/cosmoslatam/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoslatamKeeper(t)
	cosmoslatam.InitGenesis(ctx, *k, genesisState)
	got := cosmoslatam.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
