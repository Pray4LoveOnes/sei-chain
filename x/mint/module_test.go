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

package mint_test

import (
	"testing"

	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sei-protocol/sei-chain/app"
	"github.com/sei-protocol/sei-chain/x/mint"
	"github.com/sei-protocol/sei-chain/x/mint/types"
)

func TestNewProposalHandler(t *testing.T) {
	app := app.Setup(t, false, false, false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	handler := mint.NewProposalHandler(app.MintKeeper)

	newMinter := types.NewMinter(
		"2023-10-05",
		"2023-11-22",
		"usei",
		12345,
	)
	updateMinterProposal := &types.UpdateMinterProposal{
		Title:       "Test Title",
		Description: "Test Description",
		Minter:      &newMinter,
	}
	err := handler(ctx, updateMinterProposal)
	require.NoError(t, err)
	updatedMinter := app.MintKeeper.GetMinter(ctx)
	require.Equal(t, newMinter, updatedMinter)

	invalidMinter := types.NewMinter(
		"2023-11-22",
		"2023-10-05",
		"test",
		12345,
	)
	invalidProposal := &types.UpdateMinterProposal{
		Title:       "Invalid Minter",
		Description: "Invalid Minter",
		Minter:      &invalidMinter,
	}
	err = handler(ctx, invalidProposal)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "end date must be after start")
}
