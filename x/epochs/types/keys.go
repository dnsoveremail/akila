// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/akila/akila/blob/main/LICENSE)

package types

const (
	// ModuleName defines the module name
	ModuleName = "epochs"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for epochs
	RouterKey = ModuleName
)

// prefix bytes for the epochs persistent store
const (
	prefixEpoch = iota + 1
)

// KeyPrefixEpoch defines prefix key for storing epochs
var KeyPrefixEpoch = []byte{prefixEpoch}
