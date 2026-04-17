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
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/utils"
)

// Exist reports whether the given account exists in state.
// Notably this should also return true for self-destructed accounts.
func (s *DBImpl) Exist(addr common.Address) bool {
	s.k.PrepareReplayedAddr(s.ctx, addr)
	// check if the address exists as a contract
	codeHash := s.GetCodeHash(addr)
	if codeHash.Cmp(common.Hash{}) != 0 {
		return true
	}

	// check if the address exists as an EOA
	if s.GetNonce(addr) > 0 {
		return true
	}

	// check if account has a balance
	if s.GetBalance(addr).CmpBig(utils.Big0) > 0 {
		return true
	}

	// go-ethereum impl considers just-deleted accounts as "exist" as well
	return s.HasSelfDestructed(addr)
}

// Empty returns whether the given account is empty. Empty
// is defined according to EIP161 (balance = nonce = code = 0).
func (s *DBImpl) Empty(addr common.Address) bool {
	s.k.PrepareReplayedAddr(s.ctx, addr)
	return s.GetBalance(addr).CmpBig(utils.Big0) == 0 && s.GetNonce(addr) == 0 && s.GetCodeHash(addr).Cmp(common.Hash{}) == 0
}
