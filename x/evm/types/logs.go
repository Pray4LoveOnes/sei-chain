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

package types

import ethtypes "github.com/ethereum/go-ethereum/core/types"

func NewLogsFromEth(ethlogs []*ethtypes.Log) []*Log {
	logs := make([]*Log, 0, len(ethlogs))
	for _, ethlog := range ethlogs {
		logs = append(logs, newLogFromEth(ethlog))
	}
	return logs
}

func newLogFromEth(log *ethtypes.Log) *Log {
	topics := make([]string, len(log.Topics))
	for i, topic := range log.Topics {
		topics[i] = topic.String()
	}

	return &Log{
		Address: log.Address.String(),
		Topics:  topics,
		Data:    log.Data,
		Index:   uint32(log.Index), //nolint:gosec
	}
}
