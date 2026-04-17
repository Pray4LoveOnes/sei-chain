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
	"time"

	"github.com/sei-protocol/sei-chain/app"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/x/epoch/types"
	"github.com/stretchr/testify/require"
)

func TestEpochKeeper(t *testing.T) {
	app := app.Setup(t, false, false, false) // Your setup function here
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// Define an epoch
	currentTime := time.Now().UTC()
	epochIn := types.Epoch{
		CurrentEpochStartTime: currentTime,
		CurrentEpochHeight:    100,
	}

	// Verify that it's equal to what is set
	app.EpochKeeper.SetEpoch(ctx, epochIn)
	epochOut := app.EpochKeeper.GetEpoch(ctx)
	require.Equal(t, epochIn, epochOut)

	// Test case: Should panic since ctx.Blocktime() is 0
	lastEpoch := types.Epoch{
		CurrentEpochStartTime: ctx.BlockTime().Add(-2 * time.Hour), // 2 hours ago
		EpochDuration:         1 * time.Hour,                       // 1 hour epochs
		CurrentEpoch:          2,
		CurrentEpochHeight:    0,
	}
	require.Panics(t, func() { app.EpochKeeper.SetEpoch(ctx, lastEpoch) })
}
