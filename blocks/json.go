package blocks

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

// Checks if bytes is JSON.
func IsJSON(b []byte) bool {
	var s json.RawMessage
	return json.Unmarshal(b, &s) == nil
}

// Checks if string is json.
func IsJSONString(s string) bool {
	var se json.RawMessage
	return json.Unmarshal(unsafeByte(s), &se) == nil
}

func unsafeByte(s string) []byte {
	return (*[0x7fff0000]byte)(unsafe.Pointer(
		(*reflect.StringHeader)(unsafe.Pointer(&s)).Data),
	)[:len(s):len(s)]
}
