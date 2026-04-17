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

func TestAssetType(t *testing.T) {
	asset := types.Asset{AssetType: types.AssetType_TYPECW20}
	require.True(t, asset.IsCW20())
	asset.AssetType = types.AssetType_TYPECW721
	require.True(t, asset.IsCW721())
	asset.AssetType = types.AssetType_TYPENATIVE
	require.True(t, asset.IsNative())
}
