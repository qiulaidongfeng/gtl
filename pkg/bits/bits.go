// 本包提供了位操作的函数，没有使用泛型
package bits

// 获取从右起第index位的值,index>=0
func Getbit(x, index uint) uint {
	return (x >> index) & 1
}

// 设置从右起第index位的值位1,index>=0
func Setbit1(x, index uint) uint {
	return x | (1 << index)
}

// 设置从右起第index位的值位0,index>=0
func Setbit0(x, index uint) uint {
	return x & (^(1 << index))
}
