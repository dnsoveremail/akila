// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package testdata

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "akila/x/evm/types"
)

var (
	//go:embed WEVMOS.json
	WakilaJSON []byte

	// WEVMOSContract is the compiled contract of WEVMOS
	WAKILAContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WakilaJSON, &WAKILAContract)
	if err != nil {
		panic(err)
	}

	if len(WAKILAContract.Bin) == 0 {
		panic("failed to load WAKILA smart contract")
	}
}
