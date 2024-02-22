// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)

package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"akila/x/erc20/types"
	"github.com/cosmos/cosmos-sdk/codec"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// ParseRegisterCoinProposal reads and parses a ParseRegisterCoinProposal from a file.
func ParseMetadata(cdc codec.JSONCodec, metadataFile string) ([]banktypes.Metadata, error) {
	proposalMetadata := types.ProposalMetadata{}

	contents, err := os.ReadFile(filepath.Clean(metadataFile))
	if err != nil {
		return nil, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposalMetadata); err != nil {
		return nil, fmt.Errorf("failed to unmarshal proposal metadata: %w", err)
	}

	return proposalMetadata.Metadata, nil
}
