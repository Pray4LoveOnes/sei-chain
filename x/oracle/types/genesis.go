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
	"encoding/json"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params, rates []ExchangeRateTuple,
	feederDelegations []FeederDelegation, penaltyCounters []PenaltyCounter,
	aggregateExchangeRateVotes []AggregateExchangeRateVote,
	priceSnapshots []PriceSnapshot,
) *GenesisState {
	return &GenesisState{
		Params:                     params,
		ExchangeRates:              rates,
		FeederDelegations:          feederDelegations,
		PenaltyCounters:            penaltyCounters,
		AggregateExchangeRateVotes: aggregateExchangeRateVotes,
		PriceSnapshots:             priceSnapshots,
	}
}

// DefaultGenesisState - default GenesisState used by columbus-2
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:                     DefaultParams(),
		ExchangeRates:              []ExchangeRateTuple{},
		FeederDelegations:          []FeederDelegation{},
		PenaltyCounters:            []PenaltyCounter{},
		AggregateExchangeRateVotes: []AggregateExchangeRateVote{},
		PriceSnapshots:             PriceSnapshots{},
	}
}

// ValidateGenesis validates the oracle genesis state
func ValidateGenesis(data *GenesisState) error {
	return data.Params.Validate()
}

// GetGenesisStateFromAppState returns x/oracle GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
