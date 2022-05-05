package sys

func ExampleNewMmap_windows() {
	path := "文件路径"
	//文件映射起始地址离开头偏移量
	length := uint(0)
	mmap, err := NewMmap(path, length)
	//判断是否存在错误
	if err != nil {
		//错误处理
		//......
	}
	//用完后释放
	defer mmap.Close()
}
