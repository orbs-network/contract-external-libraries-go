package list

import (
	"github.com/orbs-network/contract-external-libraries-go/v1/structs"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"reflect"
)

type Serializer func(key []byte, item interface{})
type Deserializer func(key []byte) interface{}

func StringSerializer(key []byte, item interface{}) {
	state.WriteString(key, item.(string))
}

func StringDeserializer(key []byte) interface{} {
	return state.ReadString(key)
}

func BytesSerializer(key []byte, item interface{}) {
	state.WriteBytes(key, item.([]byte))
}

func BytesDeserializer(key []byte) interface{} {
	return state.ReadBytes(key)
}

func Uint64Serializer(key []byte, item interface{}) {
	state.WriteUint64(key, item.(uint64))
}

func Uint64Deserializer(key []byte) interface{} {
	return state.ReadUint64(key)
}

func Uint32Serializer(key []byte, item interface{}) {
	state.WriteUint32(key, item.(uint32))
}

func Uint32Deserializer(key []byte) interface{} {
	return state.ReadUint32(key)
}

func StructSerializer(key []byte, item interface{}) {
	structs.WriteStruct(string(key), item)
}

func StructDeserializer(t interface{}) Deserializer {
	return func (key []byte) interface{} {
		item := reflect.New(reflect.TypeOf(t)).Interface()
		err := structs.ReadStruct(string(key), item)
		if err != nil {
			panic("could not deserialize struct: " + string(key))
		}
		return item
	}
}