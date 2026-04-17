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

/*
*
Transient Module State Keys
*/
var (
	// Represents the sum of all unassociated evm account balances
	// If evm module balance is higher than this value at the end of
	// the transaction, we need to burn from module balance in order
	// for this number to align.
	GasRefundKey = []byte{0x01}
	LogsKey      = []byte{0x02}
)

/*
*
Transient Account State Keys
*/
var (
	AccountCreated = []byte{0x01}
	AccountDeleted = []byte{0x02}
)
