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
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type Logs struct {
	Ls []*ethtypes.Log `json:"logs"`
}

func (s *DBImpl) AddLog(l *ethtypes.Log) {
	l.Index = uint(len(s.GetAllLogs()))
	s.tempState.logs = append(s.tempState.logs, l)
	s.journal = append(s.journal, &addLogChange{})

	if s.logger != nil && s.logger.OnLog != nil {
		s.logger.OnLog(l)
	}
}

func (s *DBImpl) GetAllLogs() []*ethtypes.Log {
	res := make([]*ethtypes.Log, 0, len(s.tempState.logs))
	res = append(res, s.tempState.logs...)
	return res
}

func (s *DBImpl) GetLogs(common.Hash, uint64, common.Hash) []*ethtypes.Log {
	return s.GetAllLogs()
}

func (s *DBImpl) Logs() []*ethtypes.Log {
	return s.GetAllLogs()
}
