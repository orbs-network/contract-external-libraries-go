package keys

import (
	"github.com/orbs-network/contract-external-libraries-go/v1/keys"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKey(t *testing.T) {
	require.EqualValues(t, []byte("hello"), keys.Key("hello"))
	require.EqualValues(t, []byte("hello$again"), keys.Key("hello", "$", "again"))
}

