package list

import (
	"github.com/orbs-network/contract-external-libraries-go/v1/list"
	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppendOnlyList(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		l := list.NewAppendOnlyList("l")

		length := l.Append("hello")
		require.EqualValues(t, 1, length)

		item := l.Get(0)
		require.EqualValues(t, "hello", item)
	})
}
