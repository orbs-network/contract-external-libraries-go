package list

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strconv"
)

type List interface {
	Append(item interface{}) (length uint64)
	Get(index uint64) interface{}
	Length() uint64
}

func NewAppendOnlyList(name string) List {
	return &list{
		name,
	}
}

type list struct {
	name string
}

func (l *list) Append(item interface{}) (length uint64) {
	index := l.Length()
	state.WriteString(l.itemKeyName(index), item.(string))

	length = index + 1
	state.WriteUint64(l.lengthKeyName(), length)

	return
}

func (l *list) Get(index uint64) interface{} {
	return state.ReadBytes(l.itemKeyName(index))
}

func (l *list) Length() uint64 {
	return state.ReadUint64(l.lengthKeyName())
}

func (l *list) itemKeyName(index uint64) []byte {
	return []byte(l.name+"."+strconv.FormatUint(index, 10))
}

func (l *list) lengthKeyName() []byte {
	return []byte(l.name+".length")
}