// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)
package feemarket

import (
	"akila/x/feemarket/types"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for Ethermint type messages.
func NewHandler(server types.MsgServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (result *sdk.Result, err error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgUpdateParams:
			// execute state transition
			res, err := server.UpdateParams(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			err := errorsmod.Wrapf(errortypes.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, err
		}
	}
}
