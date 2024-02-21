// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package network

import (
	"encoding/json"
	"math"
	"math/big"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"akila/app"
	"akila/types"

	commonnetwork "akila/testutil/integration/common/network"
	erc20types "akila/x/erc20/types"
	evmtypes "akila/x/evm/types"
	feemarkettypes "akila/x/feemarket/types"
	infltypes "akila/x/inflation/v1/types"
	revtypes "akila/x/revenue/v1/types"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Network is the interface that wraps the methods to interact with integration test network.
//
// It was designed to avoid users to access module's keepers directly and force integration tests
// to be closer to the real user's behavior.
type Network interface {
	commonnetwork.Network

	GetEIP155ChainID() *big.Int

	// Clients
	GetERC20Client() erc20types.QueryClient
	GetEvmClient() evmtypes.QueryClient
	GetGovClient() govtypes.QueryClient
	GetRevenueClient() revtypes.QueryClient
	GetInflationClient() infltypes.QueryClient
	GetFeeMarketClient() feemarkettypes.QueryClient

	// Because to update the module params on a conventional manner governance
	// would be required, we should provide an easier way to update the params
	UpdateEvmParams(params evmtypes.Params) error
	UpdateGovParams(params govtypes.Params) error
	UpdateInflationParams(params infltypes.Params) error
	UpdateRevenueParams(params revtypes.Params) error
}

var _ Network = (*IntegrationNetwork)(nil)

// IntegrationNetwork is the implementation of the Network interface for integration tests.
type IntegrationNetwork struct {
	cfg        Config
	ctx        sdktypes.Context
	validators []stakingtypes.Validator
	app        *app.Akila

	// This is only needed for IBC chain testing setup
	valSet     *tmtypes.ValidatorSet
	valSigners map[string]tmtypes.PrivValidator
}

// New configures and initializes a new integration Network instance with
// the given configuration options. If no configuration options are provided
// it uses the default configuration.
//
// It panics if an error occurs.
func New(opts ...ConfigOption) *IntegrationNetwork {
	cfg := DefaultConfig()
	// Modify the default config with the given options
	for _, opt := range opts {
		opt(&cfg)
	}

	ctx := sdktypes.Context{}
	network := &IntegrationNetwork{
		cfg:        cfg,
		ctx:        ctx,
		validators: []stakingtypes.Validator{},
	}

	err := network.configureAndInitChain()
	if err != nil {
		panic(err)
	}
	return network
}

var (
	// bondedAmt is the amount of tokens that each validator will have initially bonded
	bondedAmt = sdktypes.TokensFromConsensusPower(1, types.PowerReduction)
	// PrefundedAccountInitialBalance is the amount of tokens that each prefunded account has at genesis
	PrefundedAccountInitialBalance = sdktypes.NewInt(int64(math.Pow10(18) * 4))
)

// configureAndInitChain initializes the network with the given configuration.
// It creates the genesis state and starts the network.
func (n *IntegrationNetwork) configureAndInitChain() error {
	// Create funded accounts based on the config and
	// create genesis accounts
	coin := sdktypes.NewCoin(n.cfg.denom, PrefundedAccountInitialBalance)
	genAccounts := createGenesisAccounts(n.cfg.preFundedAccounts)
	fundedAccountBalances := createBalances(n.cfg.preFundedAccounts, coin)

	// Create validator set with the amount of validators specified in the config
	// with the default power of 1.
	valSet, valSigners := createValidatorSetAndSigners(n.cfg.amountOfValidators)
	totalBonded := bondedAmt.Mul(sdktypes.NewInt(int64(n.cfg.amountOfValidators)))

	// Build staking type validators and delegations
	validators, err := createStakingValidators(valSet.Validators, bondedAmt)
	if err != nil {
		return err
	}

	fundedAccountBalances = addBondedModuleAccountToFundedBalances(fundedAccountBalances, sdktypes.NewCoin(n.cfg.denom, totalBonded))

	delegations := createDelegations(valSet.Validators, genAccounts[0].GetAddress())

	// Create a new EvmosApp with the following params
	akilaApp := createAkilaApp(n.cfg.chainID)

	// Configure Genesis state
	genesisState := app.NewDefaultGenesisState()

	genesisState = setAuthGenesisState(akilaApp, genesisState, genAccounts)

	stakingParams := StakingCustomGenesisState{
		denom:       n.cfg.denom,
		validators:  validators,
		delegations: delegations,
	}
	genesisState = setStakingGenesisState(akilaApp, genesisState, stakingParams)

	genesisState = setInflationGenesisState(akilaApp, genesisState)

	totalSupply := calculateTotalSupply(fundedAccountBalances)
	bankParams := BankCustomGenesisState{
		totalSupply: totalSupply,
		balances:    fundedAccountBalances,
	}
	genesisState = setBankGenesisState(akilaApp, genesisState, bankParams)

	// Init chain
	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	if err != nil {
		return err
	}

	akilaApp.InitChain(
		abcitypes.RequestInitChain{
			ChainId:         n.cfg.chainID,
			Validators:      []abcitypes.ValidatorUpdate{},
			ConsensusParams: app.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)
	// Commit genesis changes
	akilaApp.Commit()

	header := tmproto.Header{
		ChainID:            n.cfg.chainID,
		Height:             akilaApp.LastBlockHeight() + 1,
		AppHash:            akilaApp.LastCommitID().Hash,
		ValidatorsHash:     valSet.Hash(),
		NextValidatorsHash: valSet.Hash(),
		ProposerAddress:    valSet.Proposer.Address,
	}
	akilaApp.BeginBlock(abcitypes.RequestBeginBlock{Header: header})

	// Set networks global parameters
	n.app = akilaApp
	// TODO - this might not be the best way to initilize the context
	n.ctx = akilaApp.BaseApp.NewContext(false, header)
	n.validators = validators
	n.valSet = valSet
	n.valSigners = valSigners

	// Register EVMOS in denom metadata
	akilaMetadata := banktypes.Metadata{
		Description: "The native token of Akila",
		Base:        n.cfg.denom,
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    n.cfg.denom,
				Exponent: 0,
				Aliases:  []string{n.cfg.denom},
			},
			{
				Denom:    n.cfg.denom,
				Exponent: 18,
			},
		},
		Name:    "Akila",
		Symbol:  "AKILA",
		Display: n.cfg.denom,
	}
	akilaApp.BankKeeper.SetDenomMetaData(n.ctx, akilaMetadata)

	return nil
}

// GetContext returns the network's context
func (n *IntegrationNetwork) GetContext() sdktypes.Context {
	return n.ctx
}

// GetChainID returns the network's chainID
func (n *IntegrationNetwork) GetChainID() string {
	return n.cfg.chainID
}

// GetEIP155ChainID returns the network EIp-155 chainID number
func (n *IntegrationNetwork) GetEIP155ChainID() *big.Int {
	return n.cfg.eip155ChainID
}

// GetDenom returns the network's denom
func (n *IntegrationNetwork) GetDenom() string {
	return n.cfg.denom
}

// GetValidators returns the network's validators
func (n *IntegrationNetwork) GetValidators() []stakingtypes.Validator {
	return n.validators
}

// BroadcastTxSync broadcasts the given txBytes to the network and returns the response.
// TODO - this should be change to gRPC
func (n *IntegrationNetwork) BroadcastTxSync(txBytes []byte) (abcitypes.ResponseDeliverTx, error) {
	req := abcitypes.RequestDeliverTx{Tx: txBytes}
	return n.app.BaseApp.DeliverTx(req), nil
}

// Simulate simulates the given txBytes to the network and returns the simulated response.
// TODO - this should be change to gRPC
func (n *IntegrationNetwork) Simulate(txBytes []byte) (*txtypes.SimulateResponse, error) {
	gas, result, err := n.app.BaseApp.Simulate(txBytes)
	if err != nil {
		return nil, err
	}
	return &txtypes.SimulateResponse{
		GasInfo: &gas,
		Result:  result,
	}, nil
}
