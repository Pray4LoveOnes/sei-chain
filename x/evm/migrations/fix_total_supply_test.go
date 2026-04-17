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
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/migrations"
	"github.com/stretchr/testify/require"
)

func TestFixTotalSupply(t *testing.T) {
	k := &testkeeper.EVMTestApp.EvmKeeper
	ctx := testkeeper.EVMTestApp.GetContextForDeliverTx([]byte{}).WithBlockTime(time.Now())
	addr, _ := testkeeper.MockAddressPair()
	balance := sdk.NewCoins(sdk.NewCoin(sdk.MustGetBaseDenom(), sdk.OneInt()))
	k.BankKeeper().MintCoins(ctx, "evm", balance)
	k.BankKeeper().SendCoinsFromModuleToAccount(ctx, "evm", addr, balance)
	k.BankKeeper().AddWei(ctx, addr, sdk.OneInt())
	oldSupply := k.BankKeeper().GetSupply(ctx, sdk.MustGetBaseDenom()).Amount
	require.Nil(t, migrations.FixTotalSupply(ctx, k))
	require.Equal(t, oldSupply.Add(sdk.OneInt()), k.BankKeeper().GetSupply(ctx, sdk.MustGetBaseDenom()).Amount)
	require.Equal(t, sdk.OneInt(), k.BankKeeper().GetBalance(ctx, addr, sdk.MustGetBaseDenom()).Amount)
	require.Equal(t, sdk.ZeroInt(), k.BankKeeper().GetBalance(ctx, k.AccountKeeper().GetModuleAddress("evm"), sdk.MustGetBaseDenom()).Amount)
	require.Equal(t, sdk.OneInt(), k.BankKeeper().GetWeiBalance(ctx, addr))
	require.Equal(t, sdk.NewInt(999_999_999_999), k.BankKeeper().GetWeiBalance(ctx, k.AccountKeeper().GetModuleAddress("evm")))
}
