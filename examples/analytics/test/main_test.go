// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestIncrement(t *testing.T) {
	sender, _ := orbs.CreateAccount()

	h := newHarness()
	h.incrementContract.deployContract(t, sender)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := h.incrementContract.value(t, sender)
		return value == 0
	}))

	result, err := h.incrementContract.inc(t, sender)
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := h.incrementContract.value(t, sender)
		return value == 1
	}))
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

func TestIncrementWithAnalytics(t *testing.T) {
	sender, _ := orbs.CreateAccount()

	h := newHarness()
	incrementContract := h.incrementContract
	incrementContract.deployContract(t, sender)

	analyticsContract := h.analyticsContract
	analyticsContract.deployContract(t, sender)

	incrementContract.setAnalyticsContractAddress(t, sender, analyticsContract.name)

	result, err := incrementContract.inc(t, sender)
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := incrementContract.value(t, sender)
		return value == 1
	}))

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := analyticsContract.getEvents(t, sender)
		var events []Event
		if err = json.Unmarshal([]byte(value.(string)), &events); err == nil {
			return len(events) > 0 &&
				events[0].Category == "actions" &&
				events[0].Action == "increment" &&
				events[0].Label == "no label" &&
				events[0].Value == 1
		}

		return false
	}))

	result, err = incrementContract.inc(t, sender)
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)

	result, err = incrementContract.dec(t, sender)
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := analyticsContract.getAggregationByActionOverPeriodOfTime(t, sender, "actions", "count", uint64(0), uint64(0))
		aggregation := make(map[string]uint64)
		if err = json.Unmarshal([]byte(value.(string)), &aggregation); err == nil {
			return aggregation["increment"] == 2 && aggregation["decrement"] == 1
		}

		return false
	}))
}

type ERC20Event struct {
	Action string `json:"category"`
	From   string `json:"action"`
	To     string `json:"label"`
	Value  uint64 `json:"value"`

	Contract      string
	SignerAddress string
	Timestamp     uint64
}

func TestERC20WithAnalytics(t *testing.T) {
	owner, _ := orbs.CreateAccount()

	h := newHarness()
	erc20 := h.erc20Contract
	erc20.deployContract(t, owner)

	analyticsContract := h.analyticsContract
	analyticsContract.deployContract(t, owner)

	erc20.setAnalyticsContractAddress(t, owner, analyticsContract.name)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := erc20.totalSupply(t, owner)
		return value == 1000000000000000000
	}))

	firstInvestor, _ := orbs.CreateAccount()
	_, err := erc20.transfer(t, owner, firstInvestor, 1000)
	require.NoError(t, err)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := analyticsContract.getEvents(t, owner)
		var events []ERC20Event
		if err = json.Unmarshal([]byte(value.(string)), &events); err == nil {
			fmt.Println(events)
			return len(events) > 0 &&
				events[0].Action == "transfer" &&
				events[0].From == hex.EncodeToString(owner.AddressAsBytes()) &&
				events[0].To == hex.EncodeToString(firstInvestor.AddressAsBytes()) &&
				events[0].Value == 1000
		}

		return false
	}))

	for i := 0; i < 10; i++ {
		secondaryMarketParticipant, _ := orbs.CreateAccount()
		_, err := erc20.transfer(t, firstInvestor, secondaryMarketParticipant, 10+uint64(i))
		require.NoError(t, err)
	}

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := analyticsContract.getAggregationByActionOverPeriodOfTime(t, owner, "transfer", "count", uint64(0), uint64(0))
		aggregation := make(map[string]uint64)
		if err = json.Unmarshal([]byte(value.(string)), &aggregation); err == nil {
			fmt.Println(aggregation)
			return aggregation[hex.EncodeToString(owner.AddressAsBytes())] == 1 &&
				aggregation[hex.EncodeToString(firstInvestor.AddressAsBytes())] == 10
		}

		return false
	}))
}
