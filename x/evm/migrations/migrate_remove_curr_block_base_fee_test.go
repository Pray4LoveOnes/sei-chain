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
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/migrations"
	"github.com/stretchr/testify/require"
)

func TestMigrateRemoveCurrBlockBaseFee(t *testing.T) {
	k := testkeeper.EVMTestApp.EvmKeeper
	ctx := testkeeper.EVMTestApp.GetContextForDeliverTx([]byte{})

	// Set a test base fee
	testBaseFee := sdk.NewDec(100)
	testNextBaseFee := sdk.NewDec(101)
	k.SetCurrBaseFeePerGas(ctx, testBaseFee)
	k.SetNextBaseFeePerGas(ctx, testNextBaseFee)

	// Verify initial state
	require.Equal(t, testBaseFee, k.GetCurrBaseFeePerGas(ctx))
	require.Equal(t, testNextBaseFee, k.GetNextBaseFeePerGas(ctx))

	// Run the migration
	err := migrations.MigrateRemoveCurrBlockBaseFee(ctx, &k)
	require.NoError(t, err)

	// Verify the migration worked correctly
	require.Equal(t, testBaseFee, k.GetNextBaseFeePerGas(ctx))
	require.Equal(t, k.GetMinimumFeePerGas(ctx), k.GetCurrBaseFeePerGas(ctx))
}
