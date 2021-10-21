package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/daniel-farina/cosmos-latam/testutil/keeper"
	"github.com/daniel-farina/cosmos-latam/x/cosmoslatam/keeper"
	"github.com/daniel-farina/cosmos-latam/x/cosmoslatam/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmoslatamKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
