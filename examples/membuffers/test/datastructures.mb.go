// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package test

import (
	"github.com/orbs-network/membuffers/go"
	"bytes"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////
// message PhonebookEntryList

// reader

type PhonebookEntryList struct {
	// List []PhonebookEntry

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *PhonebookEntryList) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{List:%s,}", x.StringList())
}

var _PhonebookEntryList_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,}
var _PhonebookEntryList_Unions = [][]membuffers.FieldType{}

func PhonebookEntryListReader(buf []byte) *PhonebookEntryList {
	x := &PhonebookEntryList{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _PhonebookEntryList_Scheme, _PhonebookEntryList_Unions)
	return x
}

func (x *PhonebookEntryList) IsValid() bool {
	return x._message.IsValid()
}

func (x *PhonebookEntryList) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *PhonebookEntryList) Equal(y *PhonebookEntryList) bool {
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *PhonebookEntryList) ListIterator() *PhonebookEntryListListIterator {
	return &PhonebookEntryListListIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type PhonebookEntryListListIterator struct {
	iterator *membuffers.Iterator
}

func (i *PhonebookEntryListListIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *PhonebookEntryListListIterator) NextList() *PhonebookEntry {
	b, s := i.iterator.NextMessage()
	return PhonebookEntryReader(b[:s])
}

func (x *PhonebookEntryList) RawListArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *PhonebookEntryList) RawListArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *PhonebookEntryList) StringList() (res string) {
	res = "["
	for i := x.ListIterator(); i.HasNext(); {
		res += i.NextList().String() + ","
	}
	res += "]"
	return
}

// builder

type PhonebookEntryListBuilder struct {
	List []*PhonebookEntryBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *PhonebookEntryListBuilder) arrayOfList() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.List))
	for i, v := range w.List {
		res[i] = v
	}
	return res
}

func (w *PhonebookEntryListBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	err = w._builder.WriteMessageArray(buf, w.arrayOfList())
	if err != nil {
		return
	}
	return nil
}

func (w *PhonebookEntryListBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "PhonebookEntryList.List", w.arrayOfList())
	if err != nil {
		return
	}
	return nil
}

func (w *PhonebookEntryListBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *PhonebookEntryListBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *PhonebookEntryListBuilder) Build() *PhonebookEntryList {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return PhonebookEntryListReader(buf)
}

func PhonebookEntryListBuilderFromRaw(raw []byte) *PhonebookEntryListBuilder {
	return &PhonebookEntryListBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message PhonebookEntry

// reader

type PhonebookEntry struct {
	// FirstName string
	// LastName string
	// Phone uint64
	// OrbsAddress []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *PhonebookEntry) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{FirstName:%s,LastName:%s,Phone:%s,OrbsAddress:%s,}", x.StringFirstName(), x.StringLastName(), x.StringPhone(), x.StringOrbsAddress())
}

var _PhonebookEntry_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeString,membuffers.TypeUint64,membuffers.TypeBytes,}
var _PhonebookEntry_Unions = [][]membuffers.FieldType{}

func PhonebookEntryReader(buf []byte) *PhonebookEntry {
	x := &PhonebookEntry{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _PhonebookEntry_Scheme, _PhonebookEntry_Unions)
	return x
}

func (x *PhonebookEntry) IsValid() bool {
	return x._message.IsValid()
}

func (x *PhonebookEntry) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *PhonebookEntry) Equal(y *PhonebookEntry) bool {
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *PhonebookEntry) FirstName() string {
	return x._message.GetString(0)
}

func (x *PhonebookEntry) RawFirstName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *PhonebookEntry) RawFirstNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *PhonebookEntry) MutateFirstName(v string) error {
	return x._message.SetString(0, v)
}

func (x *PhonebookEntry) StringFirstName() string {
	return fmt.Sprintf(x.FirstName())
}

func (x *PhonebookEntry) LastName() string {
	return x._message.GetString(1)
}

func (x *PhonebookEntry) RawLastName() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *PhonebookEntry) RawLastNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *PhonebookEntry) MutateLastName(v string) error {
	return x._message.SetString(1, v)
}

func (x *PhonebookEntry) StringLastName() string {
	return fmt.Sprintf(x.LastName())
}

func (x *PhonebookEntry) Phone() uint64 {
	return x._message.GetUint64(2)
}

func (x *PhonebookEntry) RawPhone() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *PhonebookEntry) MutatePhone(v uint64) error {
	return x._message.SetUint64(2, v)
}

func (x *PhonebookEntry) StringPhone() string {
	return fmt.Sprintf("%x", x.Phone())
}

func (x *PhonebookEntry) OrbsAddress() []byte {
	return x._message.GetBytes(3)
}

func (x *PhonebookEntry) RawOrbsAddress() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *PhonebookEntry) RawOrbsAddressWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *PhonebookEntry) MutateOrbsAddress(v []byte) error {
	return x._message.SetBytes(3, v)
}

func (x *PhonebookEntry) StringOrbsAddress() string {
	return fmt.Sprintf("%x", x.OrbsAddress())
}

// builder

type PhonebookEntryBuilder struct {
	FirstName string
	LastName string
	Phone uint64
	OrbsAddress []byte

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *PhonebookEntryBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteString(buf, w.FirstName)
	w._builder.WriteString(buf, w.LastName)
	w._builder.WriteUint64(buf, w.Phone)
	w._builder.WriteBytes(buf, w.OrbsAddress)
	return nil
}

func (w *PhonebookEntryBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpString(prefix, offsetFromStart, "PhonebookEntry.FirstName", w.FirstName)
	w._builder.HexDumpString(prefix, offsetFromStart, "PhonebookEntry.LastName", w.LastName)
	w._builder.HexDumpUint64(prefix, offsetFromStart, "PhonebookEntry.Phone", w.Phone)
	w._builder.HexDumpBytes(prefix, offsetFromStart, "PhonebookEntry.OrbsAddress", w.OrbsAddress)
	return nil
}

func (w *PhonebookEntryBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *PhonebookEntryBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *PhonebookEntryBuilder) Build() *PhonebookEntry {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return PhonebookEntryReader(buf)
}

func PhonebookEntryBuilderFromRaw(raw []byte) *PhonebookEntryBuilder {
	return &PhonebookEntryBuilder{_overrideWithRawBuffer: raw}
}

