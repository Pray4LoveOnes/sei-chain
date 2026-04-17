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

package artifacts

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/sei-protocol/sei-chain/x/evm/artifacts/cw1155"
	"github.com/sei-protocol/sei-chain/x/evm/artifacts/cw20"
	"github.com/sei-protocol/sei-chain/x/evm/artifacts/cw721"
	"github.com/sei-protocol/sei-chain/x/evm/artifacts/native"
)

func GetParsedABI(typ string) *abi.ABI {
	switch typ {
	case "native":
		return native.GetParsedABI()
	case "cw20":
		return cw20.GetParsedABI()
	case "cw721":
		return cw721.GetParsedABI()
	case "cw1155":
		return cw1155.GetParsedABI()
	default:
		panic(fmt.Sprintf("unknown artifact type %s", typ))
	}
}

func GetBin(typ string) []byte {
	switch typ {
	case "native":
		return native.GetBin()
	case "cw20":
		return cw20.GetBin()
	case "cw721":
		return cw721.GetBin()
	case "cw1155":
		return cw1155.GetBin()
	default:
		panic(fmt.Sprintf("unknown artifact type %s", typ))
	}
}
