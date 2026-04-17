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
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/oracle/types"
)

func (k Keeper) IsVoteTarget(ctx sdk.Context, denom string) bool {
	_, err := k.GetVoteTarget(ctx, denom)
	return err == nil
}

func (k Keeper) GetVoteTargets(ctx sdk.Context) (voteTargets []string) {
	k.IterateVoteTargets(ctx, func(denom string, denomInfo types.Denom) bool {
		voteTargets = append(voteTargets, denom)
		return false
	})

	return voteTargets
}
