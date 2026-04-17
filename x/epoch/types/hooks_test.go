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

package types_test

import (
	"testing"

	"github.com/sei-protocol/sei-chain/sei-cosmos/store"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/x/epoch/keeper"
	"github.com/sei-protocol/sei-chain/x/epoch/types"
	"github.com/stretchr/testify/require"
	tmdb "github.com/tendermint/tm-db"
)

type mockEpochHooks struct {
	afterEpochEndCalled    bool
	beforeEpochStartCalled bool
	shouldPanic            bool
}

func (h *mockEpochHooks) AfterEpochEnd(_ sdk.Context, _ types.Epoch) {
	if h.shouldPanic {
		panic("AfterEpochEnd")
	}

	h.afterEpochEndCalled = true
}

func (h *mockEpochHooks) BeforeEpochStart(_ sdk.Context, _ types.Epoch) {
	if h.shouldPanic {
		panic("BeforeEpochStart")
	}

	h.beforeEpochStartCalled = true
}

func TestKeeperHooks(t *testing.T) {
	k := keeper.Keeper{}
	hooks := &mockEpochHooks{}
	k.SetHooks(hooks)

	ctx := sdk.Context{}   // setup context as required
	epoch := types.Epoch{} // setup epoch as required

	k.AfterEpochEnd(ctx, epoch)
	require.True(t, hooks.afterEpochEndCalled)

	hooks.afterEpochEndCalled = false // reset for the next test

	k.BeforeEpochStart(ctx, epoch)
	require.True(t, hooks.beforeEpochStartCalled)
}

func TestMultiHooks(t *testing.T) {
	hooks := &mockEpochHooks{}
	multiHooks := types.MultiEpochHooks{
		hooks,
	}

	db := tmdb.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ctx := sdk.NewContext(ms, tmproto.Header{}, false)
	epoch := types.Epoch{}

	multiHooks.AfterEpochEnd(ctx, epoch)
	require.True(t, hooks.afterEpochEndCalled)

	hooks.afterEpochEndCalled = false // reset for the next test

	multiHooks.BeforeEpochStart(ctx, epoch)
	require.True(t, hooks.beforeEpochStartCalled)
}

func TestMultiHooks_Panic(t *testing.T) {
	hook1 := &mockEpochHooks{shouldPanic: false}
	hook2 := &mockEpochHooks{shouldPanic: true}
	hook3 := &mockEpochHooks{shouldPanic: false}
	multiHooks := types.MultiEpochHooks{
		hook1,
		hook2,
		hook3,
	}

	db := tmdb.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ctx := sdk.NewContext(ms, tmproto.Header{}, false)
	epoch := types.Epoch{}

	multiHooks.AfterEpochEnd(ctx, epoch)
	require.True(t, hook1.afterEpochEndCalled)
	require.False(t, hook2.afterEpochEndCalled) // second hook should panic
	require.True(t, hook3.afterEpochEndCalled)  // third hook should still run after 2nd
}
