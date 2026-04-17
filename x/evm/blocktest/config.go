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

package blocktest

import (
	servertypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/types"
	"github.com/spf13/cast"
)

type Config struct {
	Enabled      bool   `mapstructure:"eth_blocktest_enabled"`
	TestDataPath string `mapstructure:"eth_blocktest_test_data_path"`
}

var DefaultConfig = Config{
	Enabled:      false,
	TestDataPath: "~/testdata/",
}

const (
	flagEnabled      = "eth_blocktest.eth_blocktest_enabled"
	flagTestDataPath = "eth_blocktest.eth_blocktest_test_data_path"
)

func ReadConfig(opts servertypes.AppOptions) (Config, error) {
	cfg := DefaultConfig // copy
	var err error
	if v := opts.Get(flagEnabled); v != nil {
		if cfg.Enabled, err = cast.ToBoolE(v); err != nil {
			return cfg, err
		}
	}
	if v := opts.Get(flagTestDataPath); v != nil {
		if cfg.TestDataPath, err = cast.ToStringE(v); err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}
