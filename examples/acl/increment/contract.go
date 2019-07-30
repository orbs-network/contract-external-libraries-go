package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/service"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(inc, dec, value, getACLContractAddress, setACLContractAddress)
var SYSTEM = sdk.Export(_init)

var COUNTER_KEY = []byte("counter")

func _init() {

}

func inc() uint64 {
	v := value() + 1
	state.WriteUint64(COUNTER_KEY, v)
	//_recordAction("increment")
	return v
}

func dec() uint64 {
	_checkPermissions("decrement")

	v := value()
	if v == 0 {
		panic("value is already 0!")
	}
	v -= 1
	state.WriteUint64(COUNTER_KEY, v)

	return v
}

func value() uint64 {
	return state.ReadUint64(COUNTER_KEY)
}

// ACL setup

var ACL_CONTRACT_ADDRESS = []byte("acl_contract_address")

func setACLContractAddress(addr string) {
	state.WriteString(ACL_CONTRACT_ADDRESS, addr)
}

func getACLContractAddress() string {
	return state.ReadString(ACL_CONTRACT_ADDRESS)
}

func _checkPermissions(action string) {
	if aclContractAddress := getACLContractAddress(); aclContractAddress != "" {
		if results := service.CallMethod(aclContractAddress,
			"checkPermissions",
			action,
		); results[0].(uint32) != 1 {
			panic("insufficient permissions for action '" + action +"'")
		}
	} else {
		panic("can't check permissions: ACL contract address is empty!")
	}
}
