//go:build go1.18
// +build go1.18

package bits

import (
	"golang.org/x/exp/constraints"
)

//获取从右起第index位的值,index>=0
func Getbit[T constraints.Integer](x, index T) T {
	return (x >> index) & 1
}

//设置从右起第index位的值位1,index>=0
func Setbit1[T constraints.Integer](x, index T) T {
	return x | (1 << index)
}

//设置从右起第index位的值位0,index>=0
func Setbit0[T constraints.Integer](x, index T) T {
	return x & (0 << index)
}
