package keeper

import (
	"github.com/meczero/mariannet/x/mariannet/types"
)

var _ types.QueryServer = Keeper{}
