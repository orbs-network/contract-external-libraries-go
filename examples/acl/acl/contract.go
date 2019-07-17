package main

import (
	"bytes"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(checkPermissions, getGuardedContractAddress, setGuardedContractAddress)
var SYSTEM = sdk.Export(_init)

func _init() {

}

func checkPermissions(action string) uint32 {
	_checkCallerContract()
	return 0
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
