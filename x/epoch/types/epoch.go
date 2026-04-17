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

import fmt "fmt"

// NewEpoch creates a new Epoch instance
func NewEpoch() Epoch {
	return Epoch{}
}

// DefaultParams returns a default set of parameters
func DefaultEpoch() Epoch {
	return NewEpoch()
}

func (e *Epoch) Validate() error {
	if e.GetGenesisTime().IsZero() {
		return fmt.Errorf("epoch genesis time cannot be zero")
	}

	if e.GetEpochDuration().Seconds() == 0 {
		return fmt.Errorf("epoch duration cannot be zero")
	}

	if e.GetGenesisTime().After(e.GetCurrentEpochStartTime()) {
		return fmt.Errorf("epoch genesis time cannot be after epoch start time")
	}

	if e.GetCurrentEpochHeight() < 0 {
		return fmt.Errorf("epoch current epoch height cannot be negative")
	}

	return nil
}
