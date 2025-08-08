package evmrpc

import (
	"github.com/spf13/viper"
)

type Config struct {
	RPCAddress string
}

func ReadConfig() Config {
	viper.SetDefault("evm_rpc.rpc_address", "localhost:8545")

	return Config{
		RPCAddress: viper.GetString("evm_rpc.rpc_address"),
	}
}
