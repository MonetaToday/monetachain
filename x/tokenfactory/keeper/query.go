package keeper

import (
	"monetachain/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
