package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func (h *erc20Contract) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	contractSource, err := ioutil.ReadFile("../erc20/contract.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.name, uint32(1), contractSource)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *erc20Contract) transfer(t *testing.T, sender *orbs.OrbsAccount, to *orbs.OrbsAccount, amount uint64) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "transfer", to.AddressAsBytes(), amount)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *erc20Contract) setAnalyticsContractAddress(t *testing.T, sender *orbs.OrbsAccount, addr string) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "setAnalyticsContractAddress", addr)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *erc20Contract) totalSupply(t *testing.T, sender *orbs.OrbsAccount) uint64 {
	query, err := h.client.CreateQuery(sender.PublicKey, h.name, "totalSupply")
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0].(uint64)
}
