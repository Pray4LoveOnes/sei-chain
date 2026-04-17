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

import (
	"gopkg.in/yaml.v2"
)

// OracleExchangeRates - array of OracleExchangeRate
type PriceSnapshots []PriceSnapshot

type (
	PriceSnapshotItems []PriceSnapshotItem
	OracleTwaps        []OracleTwap
)

// String implements fmt.Stringer interface
func (snapshots PriceSnapshots) String() string {
	out, _ := yaml.Marshal(snapshots)
	return string(out)
}

// String implements fmt.Stringer interface
func (items PriceSnapshotItems) String() string {
	out, _ := yaml.Marshal(items)
	return string(out)
}

func NewPriceSnapshotItem(denom string, exchangeRate OracleExchangeRate) PriceSnapshotItem {
	return PriceSnapshotItem{
		Denom:              denom,
		OracleExchangeRate: exchangeRate,
	}
}

func NewPriceSnapshot(priceSnapshotItems PriceSnapshotItems, snapshotTimestamp int64) PriceSnapshot {
	return PriceSnapshot{
		SnapshotTimestamp:  snapshotTimestamp,
		PriceSnapshotItems: priceSnapshotItems,
	}
}
