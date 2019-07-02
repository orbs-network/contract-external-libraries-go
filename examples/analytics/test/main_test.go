// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"encoding/json"
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
	h.deployIncrementContract(t, sender)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := h.value(t, sender)
		return value == 0
	}))

	result, err := h.inc(t, sender)
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := h.value(t, sender)
		return value == 1
	}))
}

type Event struct {
	Category string
	Action string
	Label string
	Value uint64

	Contract string
	SignerAddress string
	Timestamp uint64
}

func TestIncrementWithAnalytics(t *testing.T) {
	sender, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployIncrementContract(t, sender)
	h.deployAnalyticsContract(t, sender)

	h.setAnalyticsContractAddress(t, sender, h.analyticsContractName)

	result, err := h.inc(t, sender)
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := h.value(t, sender)
		return value == 1
	}))

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := h.getEvents(t, sender)
		var events []Event
		if err = json.Unmarshal([]byte(value.(string)), &events); err == nil {
			return len(events) > 0 &&
				events[0].Category == "action" &&
				events[0].Action == "increment" &&
				events[0].Label == "no label" &&
				events[0].Value == 1
		}

		return false
	}))
}
