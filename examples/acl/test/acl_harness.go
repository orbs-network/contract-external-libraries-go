package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func (h *aclContract) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	contractSource, err := ioutil.ReadFile("../acl/contract.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.name, uint32(1), contractSource)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *aclContract) setGuardedContractAddress(t *testing.T, sender *orbs.OrbsAccount, addr string) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "setGuardedContractAddress", addr)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *aclContract) allowAction(t *testing.T, sender *orbs.OrbsAccount, action string, address []byte) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "setPermissions", address, action, uint32(1))
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *aclContract) disallowAction(t *testing.T, sender *orbs.OrbsAccount, action string, address []byte) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.name, "setPermissions", address, action, uint32(0))
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}
