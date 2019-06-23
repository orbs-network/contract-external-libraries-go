package structs

import (
	"fmt"
	"github.com/orbs-network/contract-external-libraries-go/v1/keys"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"reflect"
)

// Takes struct as parameter:
// WriteStruct("key", Value{})
func WriteStruct(compositeKey string, item interface{}) error {
	meta := reflect.ValueOf(item)

	for i := 0; i < meta.NumField(); i++ {
		f := meta.Field(i)
		key := keys.Key(compositeKey, "$", meta.Type().Field(i).Name)

		switch v := f.Kind(); v {
		case reflect.String:
			state.WriteString(key, f.String())
		case reflect.Uint64:
			state.WriteUint64(key, f.Interface().(uint64))
		case reflect.Uint32:
			state.WriteUint32(key, f.Interface().(uint32))
		case reflect.Slice:
			state.WriteBytes(key, f.Interface().([]byte))
		default:
			return fmt.Errorf("failed to serialize key %s with type %s", key, v)
		}
	}

	return nil
}

// Takes pointer as parameter:
// ReadStruct("key", &Value{})
func ReadStruct(compositeKey string, value interface{}) error {
	meta := reflect.ValueOf(value).Elem()
	for i := 0; i < meta.NumField(); i++ {
		f := meta.Field(i)
		key := keys.Key(compositeKey, "$", meta.Type().Field(i).Name)

		fValue := meta.Field(i)

		switch v := f.Kind(); v {
		case reflect.String:
			fValue.Set(reflect.ValueOf(state.ReadString(key)))
		case reflect.Uint64:
			fValue.Set(reflect.ValueOf(state.ReadUint64(key)))
		case reflect.Uint32:
			fValue.Set(reflect.ValueOf(state.ReadUint32(key)))
		case reflect.Slice:
			bytes := state.ReadBytes(key)
			if len(bytes) > 0 { // to preserve require.EqualValues checks
				fValue.Set(reflect.ValueOf(bytes))
			}
		default:
			return fmt.Errorf("failed to deserialize key %s with type %s", key, v)
		}
	}

	return nil
}

// Takes struct as parameter:
// ClearStruct("key", Value{})
func ClearStruct(compositeKey string, value interface{}) {
	meta := reflect.ValueOf(value)
	for i := 0; i < meta.NumField(); i++ {
		key := keys.Key(compositeKey, "$"+meta.Type().Field(i).Name)
		state.Clear(key)
	}
}

// Takes struct as parameter:
// RenameStruct("oldKey", "newKey", Value{})
func RenameStruct(oldCompositeKey, newCompositeKey string, value interface{}) {
	meta := reflect.ValueOf(value)
	for i := 0; i < meta.NumField(); i++ {
		oldKey := keys.Key(oldCompositeKey, "$"+meta.Type().Field(i).Name)
		newKey := keys.Key(newCompositeKey, "$"+meta.Type().Field(i).Name)
		keys.Rename(oldKey, newKey)
	}
}
