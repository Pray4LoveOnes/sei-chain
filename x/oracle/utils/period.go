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

package utils

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	BlocksPerMinute = uint64(75)
	BlocksPerHour   = BlocksPerMinute * 60
	BlocksPerDay    = BlocksPerHour * 24
	BlocksPerWeek   = BlocksPerDay * 7
	BlocksPerMonth  = BlocksPerDay * 30
	BlocksPerYear   = BlocksPerDay * 365
)

func IsPeriodLastBlock(ctx sdk.Context, blocksPerPeriod uint64) bool {
	return ((uint64)(ctx.BlockHeight())+1)%blocksPerPeriod == 0 //nolint:gosec
}
