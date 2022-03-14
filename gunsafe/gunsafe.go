package gunsafe

import "unsafe"

//将string零拷贝转换为[]byte
func Stob(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//将[]byte零拷贝转换为string
func Btos(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
