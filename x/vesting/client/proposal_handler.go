// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)

package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"akila/x/vesting/client/cli"
)

var RegisterClawbackProposalHandler = govclient.NewProposalHandler(cli.NewClawbackProposalCmd)
