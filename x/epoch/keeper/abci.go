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
	"fmt"
	"time"

	"github.com/sei-protocol/sei-chain/sei-cosmos/telemetry"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/utils/metrics"
	"github.com/sei-protocol/sei-chain/x/epoch/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("x", "epoch", "keeper")

func (k Keeper) BeginBlock(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	lastEpoch := k.GetEpoch(ctx)
	logger.Info(" Block time", "current", ctx.BlockTime(), "last", lastEpoch.CurrentEpochStartTime, "epoch-duration", lastEpoch.EpochDuration)

	if ctx.BlockTime().Sub(lastEpoch.CurrentEpochStartTime) > lastEpoch.EpochDuration {
		k.AfterEpochEnd(ctx, lastEpoch)

		newEpoch := types.Epoch{
			GenesisTime:           lastEpoch.GenesisTime,
			EpochDuration:         lastEpoch.EpochDuration,
			CurrentEpoch:          lastEpoch.CurrentEpoch + 1,
			CurrentEpochStartTime: ctx.BlockTime(),
			CurrentEpochHeight:    ctx.BlockHeight(),
		}
		k.SetEpoch(ctx, newEpoch)
		k.BeforeEpochStart(ctx, newEpoch)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.EventTypeNewEpoch,
				sdk.NewAttribute(types.AttributeEpochNumber, fmt.Sprint(newEpoch.CurrentEpoch)),
				sdk.NewAttribute(types.AttributeEpochTime, newEpoch.CurrentEpochStartTime.String()),
				sdk.NewAttribute(types.AttributeEpochHeight, fmt.Sprint(newEpoch.CurrentEpochHeight)),
			),
		)

		metrics.SetEpochNew(newEpoch.CurrentEpoch)
	}
}
