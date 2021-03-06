package main

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/events"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/safemath/safeuint64"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/service"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(totalSupply, balanceOf, allowance, increaseAllowance, decreaseAllowance, transfer, approve, transferFrom, symbol, name, decimals,
	setAnalyticsContractAddress, getAnalyticsContractAddress)
var SYSTEM = sdk.Export(_init)
var EVENTS = sdk.Export(Approval, Transfer)

func _init() {
	state.WriteString([]byte("symbol"), "O20")
	state.WriteString([]byte("name"), "OrbsERC20Token")
	state.WriteUint32([]byte("decimals"), 18)
	_mint(address.GetSignerAddress(), 1000000000000000000)
}

//ERC20Detailed - optional parts of the EIP20 spec
func symbol() string {
	return state.ReadString([]byte("symbol"))
}

func name() string {
	return state.ReadString([]byte("name"))
}

func decimals() uint32 {
	return state.ReadUint32([]byte("decimals"))
}

// End of ERC20Detailed

// ERC20
func totalSupply() (amount uint64) {
	return state.ReadUint64([]byte("totalSupply"))
}

func balanceOf(owner []byte) uint64 {
	//validate the address
	address.ValidateAddress(owner)

	return readAccountBalance(owner)
}

func allowance(owner, spender []byte) uint64 {
	return readAccountAllowance(owner, spender)
}

func Transfer(from, to []byte, tokens uint64) {} // triggered when tokens are transferred
func transfer(to []byte, value uint64) uint32 {
	_transfer(address.GetCallerAddress(), to, value)
	return 1
}

// makes the actual data change, validates all data and arithmetic
func _transfer(from, to []byte, value uint64) {
	address.ValidateAddress(from)
	address.ValidateAddress(to)

	fromInitialBalance := readAccountBalance(from)
	newFromBalance := safeuint64.Sub(fromInitialBalance, value)
	writeAccountBalance(from, newFromBalance)

	toInitialBalance := readAccountBalance(to)
	newToBalance := safeuint64.Add(toInitialBalance, value)
	writeAccountBalance(to, newToBalance)
	events.EmitEvent(Transfer, from, to, value)
	_recordAction("transfer", hex.EncodeToString(from), hex.EncodeToString(to), value)
}

func Approval(owner, spender []byte, value uint64) {} // triggered when allowances change is called
func approve(spender []byte, value uint64) {
	address.ValidateAddress(spender)
	writeAccountAllowance(address.GetCallerAddress(), spender, value)
	events.EmitEvent(Approval, address.GetCallerAddress(), spender, value)
	_recordAction("approval", hex.EncodeToString(address.GetCallerAddress()), hex.EncodeToString(spender), value)
}

// will emit an approval event
func transferFrom(from, to []byte, value uint64) uint32 {
	_decreaseAllowance(from, address.GetCallerAddress(), value)
	_transfer(from, to, value)
	return 1
}

func increaseAllowance(spender []byte, value uint64) uint32 {
	owner := address.GetCallerAddress()
	initialAllowance := readAccountAllowance(owner, spender)
	newAllowance := safeuint64.Add(initialAllowance, value)
	writeAccountAllowance(owner, spender, newAllowance)

	events.EmitEvent(Approval, owner, spender, newAllowance)
	_recordAction("approval", hex.EncodeToString(owner), hex.EncodeToString(spender), newAllowance)
	return 1
}

func decreaseAllowance(spender []byte, value uint64) uint32 {
	_decreaseAllowance(address.GetCallerAddress(), spender, value)
	return 1
}

func _decreaseAllowance(owner, spender []byte, value uint64) {
	initialAllowance := readAccountAllowance(owner, spender)
	newAllowance := safeuint64.Sub(initialAllowance, value)
	writeAccountAllowance(owner, spender, newAllowance)

	events.EmitEvent(Approval, owner, spender, newAllowance)
	_recordAction("approval", hex.EncodeToString(owner), hex.EncodeToString(spender), newAllowance)
}

func _mint(to []byte, value uint64) {
	address.ValidateAddress(to)

	total := totalSupply()
	newTotal := safeuint64.Add(total, value)
	state.WriteUint64([]byte("totalSupply"), newTotal)

	toInitialBalance := readAccountBalance(to)
	newToBalance := safeuint64.Add(toInitialBalance, value)
	writeAccountBalance(to, newToBalance)
	events.EmitEvent(Transfer, []byte{0}, to, value)
}

// Account mapping, using a prefix for the state storage
var balancesStoragePrefix = []byte("balances.")
var allowanceStoragePrefix = []byte("allowance.")

func writeAccountBalance(owner []byte, balance uint64) {
	address.ValidateAddress(owner)
	state.WriteUint64(append(balancesStoragePrefix, owner...), balance)
}

func readAccountBalance(owner []byte) uint64 {
	address.ValidateAddress(owner)
	return state.ReadUint64(append(balancesStoragePrefix, owner...))
}

func writeAccountAllowance(owner, spender []byte, allowance uint64) {
	address.ValidateAddress(owner)
	address.ValidateAddress(spender)

	prefix := append(allowanceStoragePrefix, owner...)
	state.WriteUint64(append(prefix, spender...), allowance)
}

func readAccountAllowance(owner, spender []byte) uint64 {
	address.ValidateAddress(owner)
	address.ValidateAddress(spender)

	prefix := append(allowanceStoragePrefix, owner...)
	return state.ReadUint64(append(prefix, spender...))
}

// Analytics setup

var ANALYTICS_CONTRACT_ADDRESS = []byte("analytics_contract_address")

func setAnalyticsContractAddress(addr string) {
	state.WriteString(ANALYTICS_CONTRACT_ADDRESS, addr)
}

func getAnalyticsContractAddress() string {
	return state.ReadString(ANALYTICS_CONTRACT_ADDRESS)
}

// We re-purpose event arguments to suit our needs
func _recordAction(action string, from string, to string, value uint64) {
	if analyticsContractAddress := getAnalyticsContractAddress(); analyticsContractAddress != "" {
		service.CallMethod(analyticsContractAddress,
			"recordEvent",
			action,
			from,
			to,
			value,
		)
	}
}
