package main

import (
	"bytes"
	"encoding/hex"
	"github.com/orbs-network/contract-external-libraries-go/v1/keys"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(checkPermissions, setPermissions, getGuardedContractAddress, setGuardedContractAddress)
var SYSTEM = sdk.Export(_init)

var OWNER_KEY = []byte("owner")

func _init() {
	state.WriteBytes(OWNER_KEY, address.GetSignerAddress())
}

func checkPermissions(action string) uint32 {
	_checkCallerContract()
	return _getPermissions(address.GetSignerAddress(), action)
}

func setPermissions(address []byte, action string, level uint32) {
	_checkOwner()
	_setPermissions(address, action, level)
}

var GUARDED_CONTRACT_ADDRESS = []byte("guarded_contract_address")

func setGuardedContractAddress(addr string) {
	state.WriteString(GUARDED_CONTRACT_ADDRESS, addr)
}

func getGuardedContractAddress() string {
	return state.ReadString(GUARDED_CONTRACT_ADDRESS)
}

func _checkCallerContract() {
	if guardedContractAddress := getGuardedContractAddress(); guardedContractAddress != "" {
		if !bytes.Equal(address.GetContractAddress(guardedContractAddress), address.GetCallerAddress()) {
			panic("guarded contract address is not the same as caller address")
		}
	} else {
		panic("guarded contract address is empty!")
	}
}

func _checkOwner() {
	if !bytes.Equal(state.ReadBytes(OWNER_KEY), address.GetSignerAddress()) {
		panic("guarded contract address is not the same as caller address")
	}
}

func _setPermissions(address []byte, action string, level uint32) {
	state.WriteUint32(_permissionKey(address, action), level)
}

func _getPermissions(address []byte, action string) uint32 {
	return state.ReadUint32(_permissionKey(address, action))
}

func _permissionKey(address []byte, action string) []byte {
	return []byte(keys.Key(action, ".", hex.EncodeToString(address)))
}