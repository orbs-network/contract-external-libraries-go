package test

import (
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
	"time"
)

type harness struct {
	client                *orbs.OrbsClient
	incrementContractName string
	analyticsContractName string
}


func newHarness() *harness {
	return &harness{
		client:                orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET),
		incrementContractName: fmt.Sprintf("Inc%d", time.Now().UnixNano()),
		analyticsContractName: fmt.Sprintf("Analytics%d", time.Now().UnixNano()),
	}
}

func (h *harness) deployIncrementContract(t *testing.T, sender *orbs.OrbsAccount) {
	contractSource, err := ioutil.ReadFile("../increment/contract.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.incrementContractName, uint32(1), contractSource)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *harness) deployAnalyticsContract(t *testing.T, sender *orbs.OrbsAccount) {
	contractSource, err := ioutil.ReadFile("../analytics/contract.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.analyticsContractName, uint32(1), contractSource)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *harness) value(t *testing.T, sender *orbs.OrbsAccount) uint64 {
	query, err := h.client.CreateQuery(sender.PublicKey, h.incrementContractName, "value")
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].(uint64)
}

func (h *harness) inc(t *testing.T, sender *orbs.OrbsAccount) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.incrementContractName, "inc")
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *harness) setAnalyticsContractAddress(t *testing.T, sender *orbs.OrbsAccount, addr string) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.incrementContractName, "setAnalyticsContractAddress", addr)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *harness) recordEvent(t *testing.T, sender *orbs.OrbsAccount, eventType string, metadata string, addr string) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.analyticsContractName, "recordEvent", eventType, metadata, addr)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *harness) getEvents(t *testing.T, sender *orbs.OrbsAccount) interface{} {
	query, err := h.client.CreateQuery(sender.PublicKey, h.analyticsContractName, "getEvents")
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0]
}