// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)

package types

import (
	"errors"

	akilatypes "akila/types"
	govcdc "github.com/cosmos/cosmos-sdk/x/gov/codec"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// constants
const (
	ProposalTypeRegisterIncentive string = "RegisterIncentive"
	ProposalTypeCancelIncentive   string = "CancelIncentive"
)

// Implements Proposal Interface
var (
	_ govv1beta1.Content = &RegisterIncentiveProposal{}
	_ govv1beta1.Content = &CancelIncentiveProposal{}
)

func init() {
	govv1beta1.RegisterProposalType(ProposalTypeRegisterIncentive)
	govv1beta1.RegisterProposalType(ProposalTypeCancelIncentive)
	govcdc.ModuleCdc.Amino.RegisterConcrete(&RegisterIncentiveProposal{}, "incentives/RegisterIncentiveProposal", nil)
	govcdc.ModuleCdc.Amino.RegisterConcrete(&CancelIncentiveProposal{}, "incentives/CancelIncentiveProposal", nil)
}

// ProposalRoute returns router key for this proposal
func (*RegisterIncentiveProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*RegisterIncentiveProposal) ProposalType() string {
	return ProposalTypeRegisterIncentive
}

// ValidateBasic performs a stateless check of the proposal fields
func (rip *RegisterIncentiveProposal) ValidateBasic() error {
	return errors.New("Deprecated")
}

// ProposalRoute returns router key for this proposal
func (*CancelIncentiveProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*CancelIncentiveProposal) ProposalType() string {
	return ProposalTypeCancelIncentive
}

// ValidateBasic performs a stateless check of the proposal fields
func (rip *CancelIncentiveProposal) ValidateBasic() error {
	if err := akilatypes.ValidateAddress(rip.Contract); err != nil {
		return err
	}
	return errors.New("Deprecated")
}
