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
	analyticsContract analyticsContract
	erc20Contract     erc20Contract
}

type contract struct {
	client *orbs.OrbsClient
	name   string
}

type incrementContract contract
type analyticsContract contract
type erc20Contract contract

func newHarness() *harness {
	client := orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET)

	return &harness{
		incrementContract: incrementContract{
			name:   fmt.Sprintf("Inc%d", time.Now().UnixNano()),
			client: client,
		},
		analyticsContract: analyticsContract{
			name:   fmt.Sprintf("Analytics%d", time.Now().UnixNano()),
			client: client,
		},
		erc20Contract: erc20Contract{
			name:   fmt.Sprintf("ERC20%d", time.Now().UnixNano()),
			client: client,
		},
	}
}
