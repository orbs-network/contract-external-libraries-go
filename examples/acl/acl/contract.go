package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(checkPermissions, getGuardedContractAddress, getGuardedContractAddress)
var SYSTEM = sdk.Export(_init)

func _init() {

}

func checkPermissions(action string) uint32 {
	return 0
}


var GUARDED_CONTRACT_ADDRESS = []byte("guarded_contract_address")

func setGuardedContractAddress(addr string) {
	state.WriteString(GUARDED_CONTRACT_ADDRESS, addr)
}

func getGuardedContractAddress() string {
	return state.ReadString(GUARDED_CONTRACT_ADDRESS)
}

