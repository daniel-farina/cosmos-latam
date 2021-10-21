package keeper

import (
	"github.com/daniel-farina/cosmos-latam/x/cosmoslatam/types"
)

var _ types.QueryServer = Keeper{}
