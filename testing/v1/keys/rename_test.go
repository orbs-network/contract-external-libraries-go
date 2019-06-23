package keys

import (
	"github.com/orbs-network/contract-external-libraries-go/v1/keys"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRename(t *testing.T) {
	caller := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		state.WriteString([]byte("artist"), "David Bowie")

		require.EqualValues(t, "David Bowie", state.ReadString([]byte("artist")))
		require.Empty(t, state.ReadString([]byte("performer")))

		keys.Rename([]byte("artist"), []byte("performer"))

		require.EqualValues(t, "David Bowie", state.ReadString([]byte("performer")))
		require.Empty(t, state.ReadString([]byte("artist")))
	})
}

func TestRenameWithNonExistentKey(t *testing.T) {
	caller := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		require.NotPanics(t, func() {
			keys.Rename([]byte("artist"), []byte("performer"))
		})
	})
}