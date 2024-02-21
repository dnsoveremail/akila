// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package network

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	evmtypes "akila/x/evm/types"
	infltypes "akila/x/inflation/v1/types"
	revtypes "akila/x/revenue/v1/types"
)

func (n *IntegrationNetwork) UpdateEvmParams(params evmtypes.Params) error {
	return n.app.EvmKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateRevenueParams(params revtypes.Params) error {
	return n.app.RevenueKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateInflationParams(params infltypes.Params) error {
	return n.app.InflationKeeper.SetParams(n.ctx, params)
}

func (n *IntegrationNetwork) UpdateGovParams(params govtypes.Params) error {
	return n.app.GovKeeper.SetParams(n.ctx, params)
}