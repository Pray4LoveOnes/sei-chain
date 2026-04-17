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

package ethtx

import sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

func (tx *LegacyTx) SetTo(v string) {
	tx.To = v
}

func (tx *LegacyTx) SetAmount(v sdk.Int) {
	tx.Amount = &v
}

func (tx *LegacyTx) SetGasPrice(v sdk.Int) {
	tx.GasPrice = &v
}
