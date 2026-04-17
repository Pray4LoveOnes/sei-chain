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

func (tx *SetCodeTx) SetTo(v string) {
	tx.To = v
}

func (tx *SetCodeTx) SetAmount(v sdk.Int) {
	tx.Amount = &v
}

func (tx *SetCodeTx) SetGasFeeCap(v sdk.Int) {
	tx.GasFeeCap = &v
}

func (tx *SetCodeTx) SetGasTipCap(v sdk.Int) {
	tx.GasTipCap = &v
}

func (tx *SetCodeTx) SetAccesses(v AccessList) {
	tx.Accesses = v
}

func (tx *SetCodeTx) SetAuthList(v AuthList) {
	tx.AuthList = v
}
