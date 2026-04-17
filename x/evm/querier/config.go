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

package querier

import (
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	"github.com/spf13/cast"
)

type Config struct {
	GasLimit uint64 `mapstructure:"evm_query_gas_limit"`
}

var DefaultConfig = Config{
	GasLimit: 300000,
}

const (
	flagGasLimit = "evm_query.evm_query_gas_limit"
)

func ReadConfig(opts servertypes.AppOptions) (Config, error) {
	cfg := DefaultConfig // copy
	var err error
	if v := opts.Get(flagGasLimit); v != nil {
		if cfg.GasLimit, err = cast.ToUint64E(v); err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}
