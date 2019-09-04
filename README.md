# ORBS Network Smart Contract External Libraries

This is a set of high level libraries built on top of [Orbs Network Smart Contract SDK](https://github.com/orbs-network/orbs-contract-sdk).

## V1 APIs

### Keys

Imports from `github.com/orbs-network/contract-external-libraries-go/v1/keys`

* `Key(string, ...string)` - takes a list of strings and concatenates them into a byte array that can be used as a key for state storage.
* `Rename(oldKey, newKey)` - renames a single key in state storage.

### Structs

Imports from `github.com/orbs-network/contract-external-libraries-go/v1/structs`

**Writing and reading nested structs is not supported.**

Examples can be found in `examples/structs/contract.go`

* `WriteStruct(key, Value{})` - serializes the struct key by key. Acceptable field values are the same as the ones supported by Smart Contract SDK: `uint64`, `uint32`, `string`, and `[]byte`.
* `ReadStruct(key, &Value{})` - reads keys from state storage one by one to populate the struct that was passed as a pointer.
* `ClearStruct(key, Value{})` - removes the values of the struct from state storage.
* `RenameStruct(oldKey, newKey, Value{})` - renames state storage keys that belong to the struct.

### List

Imports from `github.com/orbs-network/contract-external-libraries-go/v1/list`

`NewAppendOnlyList(name, serializer, deserializer)` - creates a new append-only list.

List methods:

* `Append(item interface{}) (length uint64)` - appends new item
* `Get(index uint64) interface{}` - returns the item
* `Length() uint64` - returns list length
* `Iterator() Iterator` - returns an iterator for the list

`Serializer` is defined as a `func(key []byte, item interface{})`

`Deserializer` is defined as `func(key []byte) interface{}`

List of default Serializers/Deserializers:

* `StringSerializer`, `StringDeserializer`
* `Uint64Serializer`, `Uint64Deserializer`
* `Uint32Serializer`, `Uint32Deserializer`
* `BytesSerializer`, `BytesDeserializer`
* `StructSerializer`, `StructDeserializer`
