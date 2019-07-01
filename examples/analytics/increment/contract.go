package main

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/service"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(inc, value, getAnalyticsContractAddress, setAnalyticsContractAddress)
var SYSTEM = sdk.Export(_init)

var COUNTER_KEY = []byte("counter")
var ANALYTICS_CONTRACT_ADDRESS = []byte("analytics_contract_address")

func _init() {

}

func inc() uint64 {
	v := value() + 1
	state.WriteUint64(COUNTER_KEY, v)

	if analyticsContractAddress := getAnalyticsContractAddress(); analyticsContractAddress != "" {
		service.CallMethod(analyticsContractAddress,
			"recordEvent",
			"myEventType",
			"someMetadata",
			hex.EncodeToString(address.GetCallerAddress()))
	}

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
