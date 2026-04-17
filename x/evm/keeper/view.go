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

package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/artifacts"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

func (k *Keeper) QueryERCSingleOutput(ctx sdk.Context, typ string, addr common.Address, query string) (interface{}, error) {
	moduleAddr := k.AccountKeeper().GetModuleAddress(types.ModuleName)
	q, _ := artifacts.GetParsedABI(typ).Pack(query)
	r, err := k.StaticCallEVM(ctx, moduleAddr, &addr, q)
	if err != nil {
		logger.Error("Error calling address for query, skipping", "address", addr, "query", query, "err", err)
		return nil, err
	}
	o, _ := artifacts.GetParsedABI(typ).Unpack(query, r)
	if len(o) != 1 {
		logger.Error("Not getting exactly one outputs when querying address, skipping", "outputs", len(o), "address", addr, "query", query)
		return nil, err
	}
	return o[0], nil
}
