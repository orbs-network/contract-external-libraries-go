package structs

import (
	"encoding/binary"
	"github.com/orbs-network/contract-external-libraries-go/v1/structs"
	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Album struct {
	Title string
	Artist string
	Year uint64
	Artwork []byte
}

func Test_SerializeStruct(t *testing.T) {
	caller := AnAddress()

	diffs, _, _ := InServiceScope(nil, caller, func(m Mockery) {
		diamondDogs := Album{
			"Diamond Dogs",
			"David Bowie",
			1974,
			[]byte{1, 2, 3},
		}

		err := structs.WriteStruct("best-album", diamondDogs)
		require.NoError(t, err)

		value := Album{}
		err = structs.ReadStruct("best-album", &value)
		require.NoError(t, err)

		require.EqualValues(t, diamondDogs, value)
	})

	_1974 := make([]byte, 8)
	binary.LittleEndian.PutUint64(_1974, 1974)

	require.EqualValues(t, []*StateDiff{
		{[]byte("best-album$Title"), []byte("Diamond Dogs")},
		{[]byte("best-album$Artist"), []byte("David Bowie")},
		{[]byte("best-album$Year"), _1974},
		{[]byte("best-album$Artwork"), []byte{1, 2, 3}},
	}, diffs)
}

type UnserializableAlbum struct {
	Title string
	Artist string
	Year uint64
	UnserializableField interface{}
}

func Test_SerializeStructWithError(t *testing.T) {
	caller := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		diamondDogs := UnserializableAlbum{
			"Diamond Dogs",
			"David Bowie",
			1974,
			time.Now(),
		}

		err := structs.WriteStruct("best-album", diamondDogs)
		require.EqualError(t, err, "failed to serialize key best-album$UnserializableField with type interface")
	})
}

func Test_DeserializeEmptyStruct(t *testing.T) {
	caller := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		value := Album{}
		err := structs.ReadStruct("best-album", &value)
		require.NoError(t, err)

		require.EqualValues(t, Album{}, value)
	})
}

func Test_DeserializeStructWithError(t *testing.T) {
	caller := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		diamondDogs := Album{
			"Diamond Dogs",
			"David Bowie",
			1974,
			[]byte{1, 2, 3},
		}

		err := structs.WriteStruct("best-album", diamondDogs)
		require.NoError(t, err)

		value := UnserializableAlbum{}
		err = structs.ReadStruct("best-album", &value)
		require.EqualError(t, err, "failed to deserialize key best-album$UnserializableField with type interface")
	})
}

func Test_DeleteStruct(t *testing.T) {
	caller := AnAddress()

	diffs, _, _ := InServiceScope(nil, caller, func(m Mockery) {
		diamondDogs := Album{
			"Diamond Dogs",
			"David Bowie",
			1974,
			[]byte{1, 2, 3},
		}

		err := structs.WriteStruct("best-album", diamondDogs)
		require.NoError(t, err)

		structs.ClearStruct("best-album", diamondDogs)
	})

	require.EqualValues(t, []*StateDiff{
		{[]byte("best-album$Title"), []byte{}},
		{[]byte("best-album$Artist"), []byte{}},
		{[]byte("best-album$Year"), []byte{}},
		{[]byte("best-album$Artwork"), []byte{}},
	}, diffs)
}

func Test_RenameStruct(t *testing.T) {
	caller := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		diamondDogs := Album{
			"Diamond Dogs",
			"David Bowie",
			1974,
			[]byte{1, 2, 3},
		}

		err := structs.WriteStruct("best-album", diamondDogs)
		require.NoError(t, err)

		value := Album{}
		err = structs.ReadStruct("best-album", &value)
		require.NoError(t, err)
		require.EqualValues(t, diamondDogs, value)

		structs.RenameStruct("best-album", "the-best-album", value)

		oldValue := Album{}
		err = structs.ReadStruct("best-album", &oldValue)
		require.NoError(t, err)
		require.EqualValues(t, Album{}, oldValue)

		newValue := Album{}
		err = structs.ReadStruct("the-best-album", &newValue)
		require.NoError(t, err)
		require.EqualValues(t, diamondDogs, newValue)
	})
}