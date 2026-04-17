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

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/staking"
	"github.com/sei-protocol/sei-chain/x/oracle"
	"github.com/sei-protocol/sei-chain/x/oracle/keeper/testutils"
)

var (
	stakingAmt = sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction)

	randomExchangeRate        = sdk.NewDec(1700)
	anotherRandomExchangeRate = sdk.NewDecWithPrec(4882, 2) // swap rate
)

func setupWithSmallVotingPower(t *testing.T) (testutils.TestInput, sdk.Handler) {
	input := testutils.CreateTestInput(t)
	params := input.OracleKeeper.GetParams(input.Ctx)
	params.VotePeriod = 1
	params.SlashWindow = 100
	input.OracleKeeper.SetParams(input.Ctx, params)
	h := oracle.NewHandler(input.OracleKeeper)

	sh := staking.NewHandler(input.StakingKeeper)
	_, err := sh(input.Ctx, testutils.NewTestMsgCreateValidator(testutils.ValAddrs[0], testutils.ValPubKeys[0], sdk.TokensFromConsensusPower(1, sdk.DefaultPowerReduction)))
	require.NoError(t, err)

	staking.EndBlocker(input.Ctx, input.StakingKeeper)

	return input, h
}

func setup(t *testing.T) (testutils.TestInput, sdk.Handler) {
	input := testutils.CreateTestInput(t)
	params := input.OracleKeeper.GetParams(input.Ctx)
	params.VotePeriod = 1
	params.SlashWindow = 100
	input.OracleKeeper.SetParams(input.Ctx, params)

	stakingParams := input.StakingKeeper.GetParams(input.Ctx)
	stakingParams.MinCommissionRate = sdk.NewDecWithPrec(0, 2)
	input.StakingKeeper.SetParams(input.Ctx, stakingParams)

	h := oracle.NewHandler(input.OracleKeeper)

	sh := staking.NewHandler(input.StakingKeeper)

	// Validator created
	_, err := sh(input.Ctx, testutils.NewTestMsgCreateValidator(testutils.ValAddrs[0], testutils.ValPubKeys[0], stakingAmt))
	require.NoError(t, err)
	_, err = sh(input.Ctx, testutils.NewTestMsgCreateValidator(testutils.ValAddrs[1], testutils.ValPubKeys[1], stakingAmt))
	require.NoError(t, err)
	_, err = sh(input.Ctx, testutils.NewTestMsgCreateValidator(testutils.ValAddrs[2], testutils.ValPubKeys[2], stakingAmt))
	require.NoError(t, err)
	staking.EndBlocker(input.Ctx, input.StakingKeeper)

	return input, h
}

func setupN(t *testing.T, num int) (testutils.TestInput, sdk.Handler) {
	input := testutils.CreateTestInput(t)
	params := input.OracleKeeper.GetParams(input.Ctx)
	params.VotePeriod = 1
	params.SlashWindow = 100
	input.OracleKeeper.SetParams(input.Ctx, params)
	h := oracle.NewHandler(input.OracleKeeper)

	sh := staking.NewHandler(input.StakingKeeper)

	require.LessOrEqual(t, num, len(testutils.ValAddrs))

	// Validator created
	for i := 0; i < num; i++ {
		_, err := sh(input.Ctx, testutils.NewTestMsgCreateValidator(testutils.ValAddrs[i], testutils.ValPubKeys[i], stakingAmt))
		require.NoError(t, err)
	}
	staking.EndBlocker(input.Ctx, input.StakingKeeper)

	return input, h
}
