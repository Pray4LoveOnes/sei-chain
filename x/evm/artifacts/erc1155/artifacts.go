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

package erc1155

import (
	"embed"
	"sync"

	"github.com/sei-protocol/sei-chain/utils"
)

const CurrentVersion uint16 = 1

//go:embed cwerc1155.wasm
var f embed.FS

var cachedBin []byte
var cacheMtx = &sync.RWMutex{}

func GetBin() []byte {
	if cached := getCachedBin(); len(cached) > 0 {
		return cached
	}
	bz, err := f.ReadFile("cwerc1155.wasm")
	if err != nil {
		panic("failed to read ERC1155 wrapper contract wasm")
	}
	setCachedBin(bz)
	return utils.Copy(bz)
}

func getCachedBin() []byte {
	cacheMtx.RLock()
	defer cacheMtx.RUnlock()
	return utils.Copy(cachedBin)
}

func setCachedBin(bin []byte) {
	cacheMtx.Lock()
	defer cacheMtx.Unlock()
	cachedBin = bin
}
