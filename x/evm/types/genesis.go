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

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

func ValidateStream(gensisStateCh <-chan GenesisState) error {
	passedParamCheck := false
	var paramCheckErr error
	for genesisState := range gensisStateCh {
		if err := genesisState.Validate(); err != nil {
			paramCheckErr = err
		} else {
			passedParamCheck = true
		}
	}
	if !passedParamCheck {
		return paramCheckErr
	}
	return nil
}
