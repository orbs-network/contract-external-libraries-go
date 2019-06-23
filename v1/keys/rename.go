package keys

import "github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"

func Rename(oldKey []byte, newKey []byte) {
	value := state.ReadBytes(oldKey)
	state.Clear(oldKey)
	state.WriteBytes(newKey, value)
}