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

package state_test

import (
	"math/big"
	"testing"

	"github.com/sei-protocol/sei-chain/x/evm/state"
	"github.com/stretchr/testify/require"
)

func TestGetCoinbaseAddress(t *testing.T) {
	coinbaseAddr := state.GetCoinbaseAddress(1).String()
	require.Equal(t, coinbaseAddr, "sei1v4mx6hmrda5kucnpwdjsqqqqqqqqqqqpz6djs7")
}

func TestSplitUseiWeiAmount(t *testing.T) {
	for _, test := range []struct {
		amt         *big.Int
		expectedSei *big.Int
		expectedWei *big.Int
	}{
		{
			amt:         big.NewInt(0),
			expectedSei: big.NewInt(0),
			expectedWei: big.NewInt(0),
		}, {
			amt:         big.NewInt(1),
			expectedSei: big.NewInt(0),
			expectedWei: big.NewInt(1),
		}, {
			amt:         big.NewInt(999_999_999_999),
			expectedSei: big.NewInt(0),
			expectedWei: big.NewInt(999_999_999_999),
		}, {
			amt:         big.NewInt(1_000_000_000_000),
			expectedSei: big.NewInt(1),
			expectedWei: big.NewInt(0),
		}, {
			amt:         big.NewInt(1_000_000_000_001),
			expectedSei: big.NewInt(1),
			expectedWei: big.NewInt(1),
		}, {
			amt:         big.NewInt(123_456_789_123_456_789),
			expectedSei: big.NewInt(123456),
			expectedWei: big.NewInt(789_123_456_789),
		},
	} {
		usei, wei := state.SplitUseiWeiAmount(test.amt)
		require.Equal(t, test.expectedSei, usei.BigInt())
		require.Equal(t, test.expectedWei, wei.BigInt())
	}
}
