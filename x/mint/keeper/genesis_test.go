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
	"time"

	"github.com/sei-protocol/sei-chain/app"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/sei-protocol/sei-chain/testutil/nullify"

	"github.com/sei-protocol/sei-chain/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	app := app.Setup(t, false, false, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	now := time.Now()

	params := types.DefaultParams()
	params.TokenReleaseSchedule = []types.ScheduledTokenRelease{
		{
			StartDate:          now.Format(types.TokenReleaseDateFormat),
			EndDate:            now.Format(types.TokenReleaseDateFormat),
			TokenReleaseAmount: 100,
		},
	}
	genesisState := types.GenesisState{
		Params: params,
		Minter: types.Minter{
			StartDate:           now.Format(types.TokenReleaseDateFormat),
			EndDate:             now.Format(types.TokenReleaseDateFormat),
			Denom:               "usei",
			TotalMintAmount:     100,
			RemainingMintAmount: 0,
			LastMintAmount:      100,
			LastMintDate:        "2023-04-01",
			LastMintHeight:      0,
		},
	}

	app.MintKeeper.InitGenesis(ctx, &genesisState)
	got := app.MintKeeper.ExportGenesis(ctx)
	require.NotNil(t, got)
	require.Equal(t, got, &genesisState)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
