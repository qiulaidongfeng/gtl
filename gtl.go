//go:build !amd64
// +build !amd64

package gtl

import (
	"unsafe"
)

//go:linkname Memmove runtime.memmove
func Memmove(a, b unsafe.Pointer, u uint)

//go:linkname Copy gitee.com/qiulaidongfeng/gtl/internal/Copys.Copy
func Copy(dest, src unsafe.Pointer, n uint)
