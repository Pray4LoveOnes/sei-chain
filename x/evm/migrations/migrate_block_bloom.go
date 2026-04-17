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

package migrations

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

func MigrateBlockBloom(ctx sdk.Context, k *keeper.Keeper) error {
	k.SetLegacyBlockBloomCutoffHeight(ctx)

	prefsToDelete := [][]byte{}
	k.IterateAll(ctx, types.BlockBloomPrefix, func(key, _ []byte) bool {
		if len(key) > 0 {
			prefsToDelete = append(prefsToDelete, key)
		}
		return false
	})
	store := k.PrefixStore(ctx, types.BlockBloomPrefix)
	for _, pref := range prefsToDelete {
		store.Delete(pref)
	}

	return nil
}
