// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)
package miner

import (
	"akila/rpc/backend"
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/cometbft/cometbft/libs/log"
)

// API is the private miner prefixed set of APIs in the Miner JSON-RPC spec.
type API struct {
	ctx     *server.Context
	logger  log.Logger
	backend backend.EVMBackend
}

// NewPrivateAPI creates an instance of the Miner API.
func NewPrivateAPI(
	ctx *server.Context,
	backend backend.EVMBackend,
) *API {
	return &API{
		ctx:     ctx,
		logger:  ctx.Logger.With("api", "miner"),
		backend: backend,
	}
}

// SetEtherbase sets the etherbase of the miner
func (api *API) SetEtherbase(etherbase common.Address) bool {
	api.logger.Debug("miner_setEtherbase")
	return api.backend.SetEtherbase(etherbase)
}

// SetGasPrice sets the minimum accepted gas price for the miner.
func (api *API) SetGasPrice(gasPrice hexutil.Big) bool {
	api.logger.Info(api.ctx.Viper.ConfigFileUsed())
	return api.backend.SetGasPrice(gasPrice)
}
