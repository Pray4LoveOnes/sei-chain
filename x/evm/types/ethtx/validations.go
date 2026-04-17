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

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// address is validate if it is a hex string (case insensitive)
// of length 40. It may optionally have a '0x' or '0X' prefix.
func ValidateAddress(address string) error {
	if !common.IsHexAddress(address) {
		return fmt.Errorf(
			"address '%s' is not a valid ethereum hex address",
			address,
		)
	}
	return nil
}
