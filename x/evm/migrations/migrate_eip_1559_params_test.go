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

package migrations_test

import (
	"testing"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/migrations"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestMigrateEip1559Params(t *testing.T) {
	k := testkeeper.EVMTestApp.EvmKeeper
	ctx := testkeeper.EVMTestApp.NewContext(false, tmtypes.Header{})

	keeperParams := k.GetParams(ctx)
	keeperParams.BaseFeePerGas = sdk.NewDec(123)

	// Perform the migration
	err := migrations.MigrateEip1559Params(ctx, &k)
	require.NoError(t, err)

	// Ensure that the new EIP-1559 parameters were migrated and the old ones were not changed
	require.Equal(t, keeperParams.BaseFeePerGas, sdk.NewDec(123))
	require.Equal(t, keeperParams.MaxDynamicBaseFeeUpwardAdjustment, types.DefaultParams().MaxDynamicBaseFeeUpwardAdjustment)
	require.Equal(t, keeperParams.MaxDynamicBaseFeeDownwardAdjustment, types.DefaultParams().MaxDynamicBaseFeeDownwardAdjustment)
	require.Equal(t, keeperParams.TargetGasUsedPerBlock, types.DefaultParams().TargetGasUsedPerBlock)
	require.Equal(t, keeperParams.MinimumFeePerGas, types.DefaultParams().MinimumFeePerGas)
}
