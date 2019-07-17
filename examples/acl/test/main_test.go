// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
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

func TestIncrementWithACLs(t *testing.T) {
	sender, _ := orbs.CreateAccount()

	h := newHarness()
	incrementContract := h.incrementContract
	incrementContract.deployContract(t, sender)

	aclContract := h.aclContract
	aclContract.deployContract(t, sender)

	incrementContract.setACLContractAddress(t, sender, aclContract.name)

	result, err := incrementContract.inc(t, sender)
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := incrementContract.value(t, sender)
		return value == 1
	}))

	result, err = incrementContract.dec(t, sender)
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_ERROR_SMART_CONTRACT, result.ExecutionResult)
	require.EqualValues(t, "insufficient permissions for action 'decrement'", result.OutputArguments[0])
}
