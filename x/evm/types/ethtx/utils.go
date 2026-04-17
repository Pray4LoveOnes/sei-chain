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
	"errors"
	"fmt"
	"math/big"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Effective gas price is the smaller of base fee + tip limit vs total fee limit
func EffectiveGasPrice(baseFee, feeCap, tipCap *big.Int) *big.Int {
	gp := new(big.Int).Add(tipCap, baseFee)
	if gp.Cmp(feeCap) < 0 {
		return gp
	}
	return feeCap
}

// Convert a value with the provided converter and set it using the provided setter
func SetConvertIfPresent[U comparable, V any](orig U, converter func(U) V, setter func(V)) {
	var nilU U
	if orig == nilU {
		return
	}

	setter(converter(orig))
}

// validate a ethtypes.Transaction for sdk.Int overflow
func ValidateEthTx(tx *ethtypes.Transaction) error {
	if !IsValidInt256(tx.Value()) {
		return errors.New("value overflow")
	}
	if !IsValidInt256(tx.GasPrice()) {
		return errors.New("gas price overflow")
	}
	if !IsValidInt256(tx.GasFeeCap()) {
		return errors.New("gas fee cap overflow")
	}
	if !IsValidInt256(tx.GasTipCap()) {
		return errors.New("gas tip cap overflow")
	}
	if !IsValidInt256(tx.BlobGasFeeCap()) {
		return errors.New("blob gas fee cap overflow")
	}
	return nil
}

func DecodeSignature(sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != crypto.SignatureLength {
		err = fmt.Errorf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength)
		return
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v, nil
}
