// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)
package network

import (
	"akila/app"
	"akila/x/evm/statedb"
	"github.com/ethereum/go-ethereum/common"
)

// UnitTestNetwork is the implementation of the Network interface for unit tests.
// It embeds the IntegrationNetwork struct to reuse its methods and
// makes the App public for easier testing.
type UnitTestNetwork struct {
	IntegrationNetwork
	App *app.Akila
}

var _ Network = (*UnitTestNetwork)(nil)

// NewUnitTestNetwork configures and initializes a new Akila Network instance with
// the given configuration options. If no configuration options are provided
// it uses the default configuration.
//
// It panics if an error occurs.
// Note: Only uses for Unit Tests
func NewUnitTestNetwork(opts ...ConfigOption) *UnitTestNetwork {
	network := New(opts...)
	return &UnitTestNetwork{
		IntegrationNetwork: *network,
		App:                network.app,
	}
}

// GetStateDB returns the state database for the current block.
func (n *UnitTestNetwork) GetStateDB() *statedb.StateDB {
	headerHash := n.GetContext().HeaderHash()
	return statedb.New(
		n.GetContext(),
		n.App.EvmKeeper,
		statedb.NewEmptyTxConfig(common.BytesToHash(headerHash.Bytes())),
	)
}
