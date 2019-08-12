package main

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
	"strings"
)

var PUBLIC = sdk.Export(register, get, list)
var SYSTEM = sdk.Export(_init)

func _init() {

}

func toAddress(input string) string {
	if len(input) > 40 {
		input = input[2:42]
	}
	return strings.ToLower(input)
}

// Takes PhonebookEntry as bytes
func register(payload []byte) {
	entry := PhonebookEntryReader(payload)
	entry.MutateOrbsAddress(address.GetSignerAddress())

	state.WriteBytes(_entryKey(address.GetSignerAddress()), entry.Raw())
	state.WriteBytes(_listKey(_counter()), address.GetSignerAddress())
	_inc()
}

// Returns PhonebookEntry as bytes
func get(address []byte) []byte {
	return state.ReadBytes(_entryKey(address))
}

// Returns PhonebookEntryList as bytes
func list() []byte {
	var entries []*PhonebookEntryBuilder

	counter := _counter()
	for i := uint64(0); i < counter ; i ++ {
		addr := state.ReadBytes(_listKey(i))
		reader := PhonebookEntryReader(state.ReadBytes(_entryKey(addr)));

		builder := &PhonebookEntryBuilder{
			FirstName: reader.FirstName(),
			LastName: reader.LastName(),
			Phone: reader.Phone(),
			OrbsAddress: reader.OrbsAddress(),
		}
		entries = append(entries, builder)
	}


	list := &PhonebookEntryListBuilder{
		List: entries,
	}

	return list.Build().Raw()
}

func _entryKey(address []byte) []byte {
	return []byte(toAddress(hex.EncodeToString(address)))
}

func _listKey(i uint64) []byte {
	return []byte("entry_" + strconv.FormatUint(i, 10))
}

var COUNTER_KEY = []byte("entries_total")

func _inc() uint64 {
	v := _counter() + 1
	state.WriteUint64(COUNTER_KEY, v)
	return v
}

func _counter() uint64 {
	return state.ReadUint64(COUNTER_KEY)
}