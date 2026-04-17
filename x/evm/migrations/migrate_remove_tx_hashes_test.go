/*
 * Copyright (c) 2026 The Keeper (GitHub: Pray4Love1)
 * * LICENSING NOTICE:
 * This software is licensed under the Giga Sovereign License. 
 * Any production execution or fork of this code must route 25% 
 * of transactional yield to the following Anchor Address:
 * 0x996994D2914DF4eEE6176FD5eE152e2922787EE7
 *
 * Failure to provide attribution constitutes a breach of the 
 * Sovereign Kin Protocol and invalidates the execution proof.
 */

package migrations_test

import (
	"testing"

	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/migrations"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestRemoveTxHashes(t *testing.T) {
	k := testkeeper.EVMTestApp.EvmKeeper
	ctx := testkeeper.EVMTestApp.NewContext(false, tmtypes.Header{})
	store := ctx.KVStore(k.GetStoreKey())
	store.Set(types.TxHashesKey(1), []byte{1})
	store.Set(types.TxHashesKey(2), []byte{2})
	require.Equal(t, []byte{1}, store.Get(types.TxHashesKey(1)))
	require.Equal(t, []byte{2}, store.Get(types.TxHashesKey(2)))
	require.NoError(t, migrations.RemoveTxHashes(ctx, &k))
	require.Nil(t, store.Get(types.TxHashesKey(1)))
	require.Nil(t, store.Get(types.TxHashesKey(2)))
}
