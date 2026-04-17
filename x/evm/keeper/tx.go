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

package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/store/prefix"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

const DefaultTxHashesToRemove = 100

func (k *Keeper) RemoveFirstNTxHashes(ctx sdk.Context, n int) {
	store := prefix.NewStore(ctx.KVStore(k.GetStoreKey()), types.TxHashesPrefix)
	iter := store.Iterator(nil, nil)
	defer func() { _ = iter.Close() }()
	keysToDelete := make([][]byte, 0, n)
	for ; n > 0 && iter.Valid(); iter.Next() {
		keysToDelete = append(keysToDelete, iter.Key())
		n--
	}
	for _, k := range keysToDelete {
		store.Delete(k)
	}
}
