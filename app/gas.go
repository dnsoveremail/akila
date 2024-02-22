// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)

package app

import (
	"cosmossdk.io/math"
)

var (
	// MainnetMinGasPrices defines 20B aakila (or atakila) as the minimum gas price value on the fee market module.
	// See https://commonwealth.im/akila/discussion/5073-global-min-gas-price-value-for-cosmos-sdk-and-evm-transaction-choosing-a-value for reference
	MainnetMinGasPrices = math.LegacyNewDec(20_000_000_000)
	// MainnetMinGasMultiplier defines the min gas multiplier value on the fee market module.
	// 50% of the leftover gas will be refunded
	MainnetMinGasMultiplier = math.LegacyNewDecWithPrec(5, 1)
)
