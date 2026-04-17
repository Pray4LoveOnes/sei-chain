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

	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestAssociationMissingErr(t *testing.T) {
	tests := []struct {
		name            string
		address         string
		wantError       string
		wantAddressType string
	}{
		{
			name:            "EVM address",
			address:         "0x1234567890abcdef",
			wantError:       "address 0x1234567890abcdef is not linked",
			wantAddressType: "evm",
		},
		{
			name:            "SEI address",
			address:         "sei1234567890abcdef",
			wantError:       "address sei1234567890abcdef is not linked",
			wantAddressType: "sei",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := types.NewAssociationMissingErr(tt.address)

			// Test Error method
			require.Equal(t, tt.wantError, err.Error())

			// Test AddressType method
			require.Equal(t, tt.wantAddressType, err.AddressType())
		})
	}
}
