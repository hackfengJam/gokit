package cast

import (
	"reflect"
	"unsafe"
)

var (
	// dynamic data address
	firstPointer = uintptr(unsafe.Pointer(new(int)))
)

type (
	// Pointer pointer
	Pointer unsafe.Pointer
	// Bytes bytes
	Bytes []byte
	// String string
	String string
)

// invalid receiver type Pointer
// func (me Pointer) Writeable() bool {
//	addr := uintptr(me)
//	return addr >= firstPointer
// }

// Writeable writeable
func Writeable(ptr Pointer) bool {
	addr := uintptr(ptr)
	return addr >= firstPointer
}

// Writeable how to do this?
func (b Bytes) Writeable() bool {
	addr := uintptr(BytesToPointer(b))
	return addr >= firstPointer
}

// Pointer Bytes Pointer
func (b Bytes) Pointer() Pointer {
	return BytesToPointer(b)
}

// Writeable check if a string's buffer is writeable
func (s String) Writeable() bool {
	p := uintptr(StringToPointer(string(s)))
	return p >= firstPointer
}

// Pointer pointer
func (s String) Pointer() Pointer {
	return StringToPointer(string(s))
}

// StringToBytes return GoString's buffer slice(enable modify string)
func StringToBytes(s string) Bytes {
	var bh reflect.SliceHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return *(*Bytes)(unsafe.Pointer(&bh))
}

// BytesToString convert b to string without copy
func BytesToString(b []byte) String {
	return *(*String)(unsafe.Pointer(&b))
}

// StringToPointer returns &s[0], which is not allowed in go
func StringToPointer(s string) Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Pointer(p.Data)
}

// BytesToPointer returns &b[0], which is not allowed in go
func BytesToPointer(b []byte) Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return Pointer(p.Data)
}
