package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func (h *phonebookContract) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	contractSource, err := ioutil.ReadFile("../phonebook/contract.go")
	require.NoError(t, err)

	datastructuresSouce, err := ioutil.ReadFile("../phonebook/datastructures.mb.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.name, uint32(1), contractSource, datastructuresSouce)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *phonebookContract) list(t *testing.T, sender *orbs.OrbsAccount) *PhonebookEntryList {
	query, err := h.client.CreateQuery(sender.PublicKey, h.name, "list")
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return PhonebookEntryListReader(queryResponse.OutputArguments[0].([]byte))
}

func (h *phonebookContract) register(t *testing.T, sender *orbs.OrbsAccount, entry *PhonebookEntry) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "register", entry.Raw())
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *phonebookContract) get(t *testing.T, sender *orbs.OrbsAccount, address []byte) *PhonebookEntry {
	tx, err := h.client.CreateQuery(sender.PublicKey, h.name, "get", address)
	require.NoError(t, err)

	response, err := h.client.SendQuery(tx)
	require.NoError(t, err)

	return PhonebookEntryReader(response.OutputArguments[0].([]byte))
}

