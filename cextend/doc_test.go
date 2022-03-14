package cextend

func ExampleMalloc() {
	//获得内存
	ptr := Malloc(4)
	//获取成败判断
	if ptr != nil {
		//失败处理
		//......
	}
	//不用时释放获取的内存，避免内存泄露
	defer Free(ptr)
	//使用内存
	//.......
}
