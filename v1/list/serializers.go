package list

import (
	"github.com/orbs-network/contract-external-libraries-go/v1/structs"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"reflect"
)

func StringSerializer(key []byte, item interface{}) error {
	state.WriteString(key, item.(string))
	return nil
}

func StringDeserializer(key []byte) (interface{}, error) {
	return state.ReadString(key), nil
}

func StructSerializer(key []byte, item interface{}) error {
	return structs.WriteStruct(string(key), item)
}

func StructDeserializer(t interface{}) Deserializer {
	return func (key []byte) (interface{}, error) {
		item := reflect.New(reflect.TypeOf(t)).Interface()
		err := structs.ReadStruct(string(key), item)
		return item, err
	}
}