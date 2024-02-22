// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)

package testdata

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "akila/x/evm/types"
)

var (
	//go:embed VestingCaller.json
	VestingCallerJSON []byte

	// VestingCallerContract is the compiled contract calling the Vesting precompile
	VestingCallerContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(VestingCallerJSON, &VestingCallerContract)
	if err != nil {
		panic(err)
	}

	if len(VestingCallerContract.Bin) == 0 {
		panic("failed to load smart contract that calls the vesting precompile")
	}
}
