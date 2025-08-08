package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

func (k Keeper) StoreReceipt(ctx sdk.Context, receipt *types.Receipt) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetReceiptKey(receipt.TxHash)
	bz := k.cdc.MustMarshal(receipt)
	store.Set(key, bz)
}
