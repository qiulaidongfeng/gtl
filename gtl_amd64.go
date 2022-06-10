package gtl

import (
	"unsafe"

	_ "gitee.com/qiulaidongfeng/gtl/internal/Copys"
)

//go:linkname Memmove runtime.memmove
func Memmove(a, b unsafe.Pointer, u uint)

//go:linkname Copy gitee.com/qiulaidongfeng/gtl/internal/Copys.Copy_MOVSQ
func Copy(dest, src unsafe.Pointer, n uint)

//go:linkname Copy_Movups gitee.com/qiulaidongfeng/gtl/internal/Copys.Copy_SSE_Movups
func Copy_Movups(dest, src unsafe.Pointer, n uint)
