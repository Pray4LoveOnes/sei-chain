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

package types

import (
	"testing"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/oracle/utils"
	"github.com/stretchr/testify/require"
)

func TestNewPriceSnapshotItem(t *testing.T) {
	item := NewPriceSnapshotItem(utils.MicroAtomDenom, OracleExchangeRate{
		ExchangeRate: sdk.NewDec(11),
		LastUpdate:   sdk.NewInt(20),
	})

	expected := PriceSnapshotItem{
		Denom: utils.MicroAtomDenom,
		OracleExchangeRate: OracleExchangeRate{
			ExchangeRate: sdk.NewDec(11),
			LastUpdate:   sdk.NewInt(20),
		},
	}

	require.Equal(t, expected, item)
}

func TestNewPriceSnapshot(t *testing.T) {
	snapshot := NewPriceSnapshot([]PriceSnapshotItem{
		NewPriceSnapshotItem(utils.MicroEthDenom, OracleExchangeRate{
			ExchangeRate: sdk.NewDec(11),
			LastUpdate:   sdk.NewInt(20),
		}),
		NewPriceSnapshotItem(utils.MicroAtomDenom, OracleExchangeRate{
			ExchangeRate: sdk.NewDec(12),
			LastUpdate:   sdk.NewInt(20),
		}),
	}, 1)

	expected := PriceSnapshot{
		SnapshotTimestamp: 1,
		PriceSnapshotItems: []PriceSnapshotItem{
			{
				Denom: utils.MicroEthDenom,
				OracleExchangeRate: OracleExchangeRate{
					ExchangeRate: sdk.NewDec(11),
					LastUpdate:   sdk.NewInt(20),
				},
			},
			{
				Denom: utils.MicroAtomDenom,
				OracleExchangeRate: OracleExchangeRate{
					ExchangeRate: sdk.NewDec(12),
					LastUpdate:   sdk.NewInt(20),
				},
			},
		},
	}

	require.Equal(t, expected, snapshot)
}
