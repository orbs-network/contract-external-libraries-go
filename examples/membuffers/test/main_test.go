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

func TestPhonebook(t *testing.T) {
	drakulaAccount, _ := orbs.CreateAccount()

	h := newHarness()
	h.phonebookContract.deployContract(t, drakulaAccount)

	drakula, err := h.phonebookContract.register(t, drakulaAccount, (&PhonebookEntryBuilder{
		FirstName: "Count",
		LastName: "Drakula",
		Phone: 1234567890,
	}).Build())
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, drakula.ExecutionResult)

	nosferatuAccount, _ := orbs.CreateAccount()
	nosferatu, err := h.phonebookContract.register(t, nosferatuAccount, (&PhonebookEntryBuilder{
		FirstName: "Count",
		LastName: "Nosferatu",
		Phone: 987654321,
	}).Build())
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, nosferatu.ExecutionResult)


	require.True(t, test.Eventually(1*time.Second, func() bool {
		entry := h.phonebookContract.get(t, drakulaAccount, drakulaAccount.AddressAsBytes())
		return entry.FirstName() == "Count" && entry.LastName() == "Drakula" && entry.Phone() == uint64(1234567890)
	}))

	require.True(t, test.Eventually(1*time.Second, func() bool {
		list  := h.phonebookContract.list(t, drakulaAccount)
		var entries []*PhonebookEntry
		for i := list.ListIterator(); i.HasNext(); {
			entries = append(entries, i.NextList())
		}

		return len(entries) == 2 && entries[0].LastName() == "Drakula" && entries[1].LastName() == "Nosferatu"
	}))
}

