package keeper

import (
	"github.com/meczero/mariannet/x/marian/types"
)

var _ types.QueryServer = Keeper{}
