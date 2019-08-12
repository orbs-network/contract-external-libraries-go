package test

import (
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"time"
)

type harness struct {
	phonebookContract phonebookContract
}

type contract struct {
	client *orbs.OrbsClient
	name   string
}

type phonebookContract contract

func newHarness() *harness {
	client := orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET)

	return &harness{
		phonebookContract: phonebookContract{
			name:   fmt.Sprintf("Phone%d", time.Now().UnixNano()),
			client: client,
		},
	}
}
