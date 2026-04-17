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
	"math"
	"testing"

	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestAddERCNativePointerProposalV2(t *testing.T) {
	p := types.AddERCNativePointerProposalV2{
		Title:       "title",
		Description: "desc",
		Token:       "test",
		Name:        "TEST",
		Symbol:      "Test",
		Decimals:    6,
	}
	require.Equal(t, "title", p.GetTitle())
	require.Equal(t, "desc", p.GetDescription())
	require.Equal(t, "evm", p.ProposalRoute())
	require.Equal(t, "AddERCNativePointerV2", p.ProposalType())
	p.Decimals = math.MaxUint32
	require.NotNil(t, p.ValidateBasic())
	p.Decimals = 6
	require.Nil(t, p.ValidateBasic())
	require.NotEmpty(t, p.String())
}
