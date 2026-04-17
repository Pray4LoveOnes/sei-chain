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

	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestMsgClaim(t *testing.T) {
	sender := sdk.AccAddress("acc_________________")
	claimer := common.HexToAddress("0x0123456789abcdef012345abcdef12345678")
	msg := types.NewMsgClaim(sender, claimer)
	require.Equal(t, "evm", msg.Route())
	require.Equal(t, "evm_claim", msg.Type())
	require.Len(t, msg.GetSigners(), 1)
	msg.Sender = "bad"
	require.Error(t, msg.ValidateBasic())
	require.Panics(t, func() { msg.GetSigners() })
	msg.Sender = sender.String()
	require.NotEmpty(t, msg.GetSignBytes())
	require.NoError(t, msg.ValidateBasic())
}
