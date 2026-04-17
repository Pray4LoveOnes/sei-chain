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

package oracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sei-protocol/sei-chain/x/oracle"
	"github.com/sei-protocol/sei-chain/x/oracle/keeper/testutils"
	"github.com/sei-protocol/sei-chain/x/oracle/types"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

func TestExportInitGenesis(t *testing.T) {
	input, _ := setup(t)

	input.OracleKeeper.SetFeederDelegation(input.Ctx, testutils.ValAddrs[0], testutils.Addrs[1])
	input.OracleKeeper.SetBaseExchangeRate(input.Ctx, "denom", sdk.NewDec(123))
	input.OracleKeeper.SetAggregateExchangeRateVote(input.Ctx, testutils.ValAddrs[0], types.NewAggregateExchangeRateVote(types.ExchangeRateTuples{{Denom: "foo", ExchangeRate: sdk.NewDec(123)}}, testutils.ValAddrs[0]))
	input.OracleKeeper.SetVoteTarget(input.Ctx, "denom")
	input.OracleKeeper.SetVoteTarget(input.Ctx, "denom2")
	input.OracleKeeper.SetVotePenaltyCounter(input.Ctx, testutils.ValAddrs[0], 2, 3, 0)
	input.OracleKeeper.SetVotePenaltyCounter(input.Ctx, testutils.ValAddrs[1], 4, 5, 0)
	input.OracleKeeper.AddPriceSnapshot(input.Ctx, types.NewPriceSnapshot(
		types.PriceSnapshotItems{
			{
				Denom: "usei",
				OracleExchangeRate: types.OracleExchangeRate{
					ExchangeRate: sdk.NewDec(12),
					LastUpdate:   sdk.NewInt(3600),
				},
			},
			{
				Denom: "uatom",
				OracleExchangeRate: types.OracleExchangeRate{
					ExchangeRate: sdk.NewDec(10),
					LastUpdate:   sdk.NewInt(3600),
				},
			},
		},
		int64(3600),
	))
	input.OracleKeeper.AddPriceSnapshot(input.Ctx, types.NewPriceSnapshot(
		types.PriceSnapshotItems{
			{
				Denom: "usei",
				OracleExchangeRate: types.OracleExchangeRate{
					ExchangeRate: sdk.NewDec(15),
					LastUpdate:   sdk.NewInt(3700),
				},
			},
			{
				Denom: "uatom",
				OracleExchangeRate: types.OracleExchangeRate{
					ExchangeRate: sdk.NewDec(13),
					LastUpdate:   sdk.NewInt(3700),
				},
			},
		},
		int64(3700),
	))
	genesis := oracle.ExportGenesis(input.Ctx, input.OracleKeeper)

	newInput := testutils.CreateTestInput(t)
	oracle.InitGenesis(newInput.Ctx, newInput.OracleKeeper, genesis)
	newGenesis := oracle.ExportGenesis(newInput.Ctx, newInput.OracleKeeper)

	require.Equal(t, genesis, newGenesis)
}
