package list

import (
	"github.com/orbs-network/contract-external-libraries-go/v1/list"
	"github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
)

type Album struct {
	Title string
	Artist string
	Year uint64
	Artwork []byte
}

func TestAppendOnlyList(t *testing.T) {
	t.Run("with string serializer", func(t *testing.T) {
		unit.InServiceScope(nil, nil, func(m unit.Mockery) {
			l := list.NewAppendOnlyList("l", list.StringSerializer, list.StringDeserializer)

			length := l.Append("hello")
			require.EqualValues(t, 1, length)

			item := l.Get(0)
			require.EqualValues(t, "hello", item)

			l.Append("world")
			require.EqualValues(t, l.Get(1), "world")
		})
	})

	t.Run("with struct serializer", func(t *testing.T) {
		unit.InServiceScope(nil, nil, func(m unit.Mockery) {
			l := list.NewAppendOnlyList("l", list.StructSerializer, list.StructDeserializer(Album{}))

			diamondDogs := Album{
				Title: "Diamond Dogs",
				Artist: "David Bowie",
				Year: 1974,
				Artwork: []byte{0, 1, 2, 3},
			}
			length := l.Append(diamondDogs)
			require.EqualValues(t, 1, length)

			item := l.Get(0)
			require.EqualValues(t, &diamondDogs, item)
		})
	})

	t.Run("with iterator", func(t *testing.T) {
		unit.InServiceScope(nil, nil, func(m unit.Mockery) {
			l := list.NewAppendOnlyList("l", list.StringSerializer, list.StringDeserializer)

			l.Append("Diamond Dogs")
			l.Append("by")
			l.Append("David Bowie")

			var list []string
			for i := l.Iterator(); i.Next(); {
				list = append(list, i.Value().(string))
			}

			require.EqualValues(t, []string{"Diamond Dogs", "by", "David Bowie"}, list)
		})
	})
}
