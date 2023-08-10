package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/meczero/mariannet/x/marian/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func (k Keeper) CreateMarian(ctx sdk.Context, msg types.MsgCreateMarian) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(msg.Amount)
	store.Set([]byte(msg.Creator.String()), bz)
}

func (k Keeper) GetMarian(ctx sdk.Context, creator sdk.AccAddress) int64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(creator.String())) {
		return 0
	}
	bz := store.Get([]byte(creator.String()))
	var amount int64
	k.cdc.MustUnmarshalBinaryBare(bz, &amount)
	return amount
}
