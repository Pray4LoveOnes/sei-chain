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

package state

import (
	"encoding/binary"
	"math/big"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// UseiToSweiMultiplier Fields that were denominated in usei will be converted to swei (1usei = 10^12swei)
// for existing Ethereum application (which assumes 18 decimal points) to display properly.
var UseiToSweiMultiplier = big.NewInt(1_000_000_000_000)
var SdkUseiToSweiMultiplier = sdk.NewIntFromBigInt(UseiToSweiMultiplier)

var CoinbaseAddressPrefix = []byte("evm_coinbase")

func GetCoinbaseAddress(txIdx int) sdk.AccAddress {
	txIndexBz := make([]byte, 8)
	binary.BigEndian.PutUint64(txIndexBz, uint64(txIdx)) //nolint:gosec
	return append(CoinbaseAddressPrefix, txIndexBz...)
}

func SplitUseiWeiAmount(amt *big.Int) (sdk.Int, sdk.Int) {
	wei := new(big.Int).Mod(amt, UseiToSweiMultiplier)
	usei := new(big.Int).Quo(amt, UseiToSweiMultiplier)
	return sdk.NewIntFromBigInt(usei), sdk.NewIntFromBigInt(wei)
}
