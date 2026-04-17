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

package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

func TransferWithoutEvents(db vm.StateDB, sender, recipient common.Address, amount *uint256.Int) {
	sdb := GetDBImpl(db)
	if sdb == nil {
		panic("EventlessTransfer only works with DBImpl")
	}
	sdb.DisableEvents()
	defer sdb.EnableEvents()

	sdb.SubBalance(sender, amount, tracing.BalanceChangeTransfer)
	sdb.AddBalance(recipient, amount, tracing.BalanceChangeTransfer)
}
