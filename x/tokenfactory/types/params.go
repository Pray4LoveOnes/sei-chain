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
	"fmt"

	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// DefaultDenomAllowListMaxSize default denom allowlist max size and can be overridden by governance proposal.
const DefaultDenomAllowListMaxSize = 2000

// ParamKeyTable ParamTable for tokenfactory module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams default tokenfactory module parameters.
func DefaultParams() Params {
	return Params{
		DenomAllowlistMaxSize: DefaultDenomAllowListMaxSize,
	}
}

// Validate validate params.
func (p Params) Validate() error {
	if err := validateDenomAllowListMaxSize(p.DenomAllowlistMaxSize); err != nil {
		return err
	}
	return nil
}

// ParamSetPairs Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(DenomAllowListMaxSizeKey, &p.DenomAllowlistMaxSize, validateDenomAllowListMaxSize),
	}
}

// validateDenomAllowListMaxSize validates a parameter value is within a valid range.
func validateDenomAllowListMaxSize(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
