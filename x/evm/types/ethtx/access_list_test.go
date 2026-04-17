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

package ethtx

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestAccessList(t *testing.T) {
	require.Nil(t, NewAccessList(nil))
	ethAccessList := mockAccessList()
	require.Equal(t, ethAccessList, *NewAccessList(&ethAccessList).ToEthAccessList())
}

func mockAccessList() ethtypes.AccessList {
	return ethtypes.AccessList{
		ethtypes.AccessTuple{
			Address:     common.Address{'a'},
			StorageKeys: []common.Hash{{'b'}},
		},
	}
}
