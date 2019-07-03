package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/service"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(inc, dec, value, getAnalyticsContractAddress, setAnalyticsContractAddress)
var SYSTEM = sdk.Export(_init)

var COUNTER_KEY = []byte("counter")
var ANALYTICS_CONTRACT_ADDRESS = []byte("analytics_contract_address")

func _init() {

}

func inc() uint64 {
	v := value() + 1
	state.WriteUint64(COUNTER_KEY, v)
	_recordAction("increment")

	return v
}

func dec() uint64 {
	v := value()
	if v == 0 {
		panic("value is already 0!")
	}
	v -= 1
	state.WriteUint64(COUNTER_KEY, v)
	_recordAction("decrement")

	return v
}

func value() uint64 {
	return state.ReadUint64(COUNTER_KEY)
}

func setAnalyticsContractAddress(addr string) {
	state.WriteString(ANALYTICS_CONTRACT_ADDRESS, addr)
}

func getAnalyticsContractAddress() string {
	return state.ReadString(ANALYTICS_CONTRACT_ADDRESS)
}

func _recordAction(action string) {
	if analyticsContractAddress := getAnalyticsContractAddress(); analyticsContractAddress != "" {
		service.CallMethod(analyticsContractAddress,
			"recordEvent",
			"actions",  // category
			action,     // action
			"no label", // label
			uint64(1),  // value
		)
	}
}
