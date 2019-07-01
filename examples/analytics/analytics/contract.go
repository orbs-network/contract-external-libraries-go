package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/orbs-network/contract-external-libraries-go/v1/keys"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/env"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
	"strings"
)

var PUBLIC = sdk.Export(recordEvent, getEvents)
var SYSTEM = sdk.Export(_init)

var COUNTER_KEY = []byte("counter")

func _init() {

}

type Event struct {
	Type string
	Metadata string

	Contract string
	SignerAddress string
	Timestamp uint64
}

func recordEvent(eventType string, metadata string, addr string) {
	event := Event{
		Type:          eventType,
		Metadata:      metadata,
		SignerAddress: _toAddress(addr),
		Contract:      hex.EncodeToString(address.GetCallerAddress()),
		Timestamp:     env.GetBlockTimestamp(),
	}
	//state.WriteString([]byte("hello"), "world")

	_saveEvent(event)
}

func getEvents() string {
	var events []Event

	events_total := _value()
	for i := uint64(0); i < events_total; i++ {
		//event := Event{}
		//structs.ReadStruct("events_" + strconv.FormatUint(_value(), 10), &event)
		events = append(events, _getEvent(i))
	}

	rawJson, _ := json.Marshal(events)
	return string(rawJson)
}

func _toAddress(input string) string {
	if len(input) > 40 {
		input = input[2:42]
	}
	return strings.ToLower(input)
}

func _inc() uint64 {
	v := _value() + 1
	state.WriteUint64(COUNTER_KEY, v)
	return v
}

func _value() uint64 {
	return state.ReadUint64(COUNTER_KEY)
}

func _saveEvent(event Event) {
	id := strconv.FormatUint(_value(), 10)

	state.WriteString(keys.Key("events_type", id), event.Type)
	state.WriteString(keys.Key("events_metadata", id), event.Metadata)
	state.WriteString(keys.Key("events_signer", id), event.SignerAddress)
	state.WriteString(keys.Key("events_contract", id), event.Contract)
	state.WriteUint64(keys.Key("events_timestamp", id), event.Timestamp)

	_inc()
}

func _getEvent(i uint64) Event {
	id := strconv.FormatUint(i, 10)

	return Event{
		Type: state.ReadString(keys.Key("events_type", id)),
		Metadata: state.ReadString(keys.Key("events_metadata", id)),
		SignerAddress: state.ReadString(keys.Key("events_signer", id)),
		Contract: state.ReadString(keys.Key("events_contract", id)),
		Timestamp: state.ReadUint64(keys.Key("events_timestamp", id)),
	}
}