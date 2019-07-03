package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/orbs-network/contract-external-libraries-go/v1/structs"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/env"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
)

var PUBLIC = sdk.Export(recordEvent, getEvents, getAggregationByActionOverPeriodOfTime)
var SYSTEM = sdk.Export(_init)

var COUNTER_KEY = []byte("counter")

func _init() {

}

type Event struct {
	Category string
	Action   string
	Label    string
	Value    uint64

	Contract      string
	SignerAddress string
	Timestamp     uint64
}

// Required: category, action
// Optional: label, value, metadata
// Metadata should be in JSON format
func recordEvent(eventCategory string, eventAction string, eventLabel string, eventValue uint64) {
	event := Event{
		Category:      eventCategory,
		Action:        eventAction,
		Label:         eventLabel,
		Value:         eventValue,
		SignerAddress: hex.EncodeToString(address.GetSignerAddress()),
		Contract:      hex.EncodeToString(address.GetCallerAddress()),
		Timestamp:     env.GetBlockTimestamp(),
	}

	structs.WriteStruct("events_"+strconv.FormatUint(_value(), 10), event)
	_inc()
}

// This is just an example, we should be able to retrieve raw data with scrolling interface
func getEvents() string {
	var events []Event

	events_total := _value()
	for i := uint64(0); i < events_total; i++ {
		event := Event{}
		structs.ReadStruct("events_"+strconv.FormatUint(i, 10), &event)
		events = append(events, event)
	}

	rawJson, _ := json.Marshal(events)
	return string(rawJson)
}

// aggType - count, sum, min, max
// This is just an example, real aggregation should NOT be done in a smart contract
func getAggregationByActionOverPeriodOfTime(eventCategory string, aggregationType string, startTime uint64, endTime uint64) string {
	aggregation := make(map[string]uint64)

	events_total := _value()
	for i := uint64(0); i < events_total; i++ {
		event := Event{}
		structs.ReadStruct("events_"+strconv.FormatUint(i, 10), &event)

		if startTime != 0 && event.Timestamp < startTime {
			continue
		}
		if endTime != 0 && event.Timestamp > endTime {
			continue
		}
		if eventCategory != "" && event.Category != eventCategory {
			continue
		}

		v := aggregation[event.Action]
		switch aggregationType {
		case "count":
			v += 1
		case "sum":
			v += event.Value
		case "max":
			v = _max(v, event.Value)
		case "min":
			if i == 0 {
				v = event.Value
			} else {
				v = _min(v, event.Value)
			}
		}

		aggregation[event.Action] = v
	}

	rawJson, _ := json.Marshal(aggregation)
	return string(rawJson)
}

func _inc() uint64 {
	v := _value() + 1
	state.WriteUint64(COUNTER_KEY, v)
	return v
}

func _value() uint64 {
	return state.ReadUint64(COUNTER_KEY)
}

func _max(x, y uint64) uint64 {
	if x < y {
		return y
	}
	return x
}

func _min(x, y uint64) uint64 {
	if x > y {
		return y
	}
	return x
}
