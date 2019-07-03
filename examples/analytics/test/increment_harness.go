package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func (h *incrementContract) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	contractSource, err := ioutil.ReadFile("../increment/contract.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.name, uint32(1), contractSource)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *analyticsContract) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	contractSource, err := ioutil.ReadFile("../analytics/contract.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.name, uint32(1), contractSource)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *incrementContract) value(t *testing.T, sender *orbs.OrbsAccount) uint64 {
	query, err := h.client.CreateQuery(sender.PublicKey, h.name, "value")
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].(uint64)
}

func (h *incrementContract) inc(t *testing.T, sender *orbs.OrbsAccount) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "inc")
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *incrementContract) dec(t *testing.T, sender *orbs.OrbsAccount) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "dec")
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *incrementContract) setAnalyticsContractAddress(t *testing.T, sender *orbs.OrbsAccount, addr string) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "setAnalyticsContractAddress", addr)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

