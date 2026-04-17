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

package types

import "github.com/ethereum/go-ethereum/common"

func (w *Whitelist) IsHashInWhiteList(h common.Hash) bool {
	for _, s := range w.Hashes {
		if s == h.Hex() {
			return true
		}
	}
	return false
}
