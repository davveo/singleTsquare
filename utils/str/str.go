package str

import "unsafe"

func StrToByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func ByteTostr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
