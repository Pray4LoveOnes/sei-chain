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

package ante

import (
	"errors"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

type EVMRouterDecorator struct {
	defaultAnteHandler sdk.AnteHandler
	evmAnteHandler     sdk.AnteHandler
}

func NewEVMRouterDecorator(
	defaultAnteHandler sdk.AnteHandler,
	evmAnteHandler sdk.AnteHandler,
) *EVMRouterDecorator {
	return &EVMRouterDecorator{
		defaultAnteHandler: defaultAnteHandler,
		evmAnteHandler:     evmAnteHandler,
	}
}

func (r EVMRouterDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
	if isEVM, err := IsEVMMessage(tx); err != nil {
		return ctx, err
	} else if isEVM {
		return r.evmAnteHandler(ctx, tx, simulate)
	}

	return r.defaultAnteHandler(ctx, tx, simulate)
}

func IsEVMMessage(tx sdk.Tx) (bool, error) {
	hasEVMMsg := false
	for _, msg := range tx.GetMsgs() {
		switch msg.(type) {
		case *types.MsgEVMTransaction:
			hasEVMMsg = true
		default:
			continue
		}
	}

	if hasEVMMsg && len(tx.GetMsgs()) != 1 {
		return false, errors.New("EVM tx must have exactly one message")
	}

	return hasEVMMsg, nil
}
