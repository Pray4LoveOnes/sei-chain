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

	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/stretchr/testify/require"
)

func TestGetFeeCollectorAddress(t *testing.T) {
	k := &testkeeper.EVMTestApp.EvmKeeper
	ctx := testkeeper.EVMTestApp.GetContextForDeliverTx([]byte{})
	addr, err := k.GetFeeCollectorAddress(ctx)
	require.Nil(t, err)
	expected := k.GetEVMAddressOrDefault(ctx, k.AccountKeeper().GetModuleAddress("fee_collector"))
	require.Equal(t, expected.Hex(), addr.Hex())
}

func TestGetCoinbaseAddress(t *testing.T) {
	require.Equal(t, "0x27F7B8B8B5A4e71E8E9aA671f4e4031E3773303F", keeper.GetCoinbaseAddress().Hex())
}
