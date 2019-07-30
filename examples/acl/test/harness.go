package test

import (
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"time"
)

type harness struct {
	incrementContract incrementContract
	aclContract aclContract
}

type contract struct {
	client *orbs.OrbsClient
	name   string
}

type incrementContract contract
type aclContract contract

func newHarness() *harness {
	client := orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET)

	return &harness{
		incrementContract: incrementContract{
			name:   fmt.Sprintf("Inc%d", time.Now().UnixNano()),
			client: client,
		},
		aclContract: aclContract{
			name:   fmt.Sprintf("ACL%d", time.Now().UnixNano()),
			client: client,
		},
	}
}
