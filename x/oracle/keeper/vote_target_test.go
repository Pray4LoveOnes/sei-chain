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

package keeper_test

import (
	"testing"

	"github.com/sei-protocol/sei-chain/x/oracle/keeper/testutils"
	"github.com/stretchr/testify/require"
)

func TestKeeper_GetVoteTargets(t *testing.T) {
	input := testutils.CreateTestInput(t)

	input.OracleKeeper.ClearVoteTargets(input.Ctx)

	expectedTargets := []string{"bar", "foo", "whoowhoo"}
	for _, target := range expectedTargets {
		input.OracleKeeper.SetVoteTarget(input.Ctx, target)
	}

	targets := input.OracleKeeper.GetVoteTargets(input.Ctx)
	require.Equal(t, expectedTargets, targets)
}

func TestKeeper_IsVoteTarget(t *testing.T) {
	input := testutils.CreateTestInput(t)

	input.OracleKeeper.ClearVoteTargets(input.Ctx)

	validTargets := []string{"bar", "foo", "whoowhoo"}
	for _, target := range validTargets {
		input.OracleKeeper.SetVoteTarget(input.Ctx, target)
		require.True(t, input.OracleKeeper.IsVoteTarget(input.Ctx, target))
	}
}
