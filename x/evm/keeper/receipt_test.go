package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestStoreReceipt(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	receipt := &types.Receipt{TxHash: []byte("abc123")}

	keeper.StoreReceipt(ctx, receipt)
	store := ctx.KVStore(keeper.StoreKey())
	bz := store.Get(types.GetReceiptKey(receipt.TxHash))

	require.NotNil(t, bz)
}
