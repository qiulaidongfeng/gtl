package sys

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

const (
	//页面只读
	PAGE_READ = windows.PAGE_READONLY
	//页面可执行
	PAGE_EXEC = windows.PAGE_EXECUTE
	//页面读写
	PAGE_RW = windows.PAGE_READWRITE
	//页面可读，可写，可执行
	PAGE_WX = windows.PAGE_EXECUTE_READ
	//页面可读，可写，可执行
	PAGE_RWX = PAGE_RW
	//写复制
	PAGE_WRITECOPY = windows.PAGE_EXECUTE_WRITECOPY
)

const (
	//写复制
	FILE_MAP_COPY = windows.FILE_MAP_COPY
	//读写
	FILE_MAP_WRITE = windows.FILE_MAP_WRITE
	//只读
	FILE_MAP_READ = windows.FILE_MAP_READ
	//可执行
	FILE_MAP_EXECUTE = windows.FILE_MAP_EXECUTE
	//可读，可写，可执行
	FILE_MAP_RWX = FILE_MAP_WRITE | FILE_MAP_EXECUTE
)

var (
	//系统内存页的尺寸
	Pagesize = uint(os.Getpagesize())
)

//内存映射的结构体
type Mmap struct {
	file       *os.File
	mmaphandle windows.Handle
	length     uint
	addr       uintptr
}

//以读写模式打开文件，0777权限位，可读可写
func NewMmap(path string, length uint) (m *Mmap, err error) {
	m, err = NewMmapAll(path, os.O_RDWR|os.O_CREATE, 0777, length, PAGE_RWX, FILE_MAP_WRITE)
	return
	/*//打开文件
	m.file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	//改变文件大小为length
	err = m.file.Truncate(int64(length))
	if err != nil {
		return nil, err
	}
	m.length = length
	//进行系统调用实现共享内存
	m.mmaphandle, err = windows.CreateFileMapping(
		windows.Handle(m.file.Fd()),
		nil,
		PAGE_RWX,
		uint32(length>>32),
		uint32(length),
		nil)
	if err != nil {
		return nil, err
	}
	//使文件大小等于内存页倍数
	pagelen := length%uint(Pagesize) + 1
	size := pagelen * Pagesize
	m.addr, err = windows.MapViewOfFile(m.mmaphandle, FILE_MAP_RWX, 0, 0, uintptr(size))
	if err != nil {
		return nil, err
	}
	return m, nil
	*/
}

//以自定义模式与自定义权限位打开文件，自定义是否读写执行
func NewMmapAll(path string, osflag int, perm os.FileMode, length uint, prot uint32, fileflag uint32) (m *Mmap, err error) {
	var m1 *Mmap = new(Mmap)
	//打开文件
	m1.file, err = os.OpenFile(path, osflag, perm)
	if err != nil {
		return nil, err
	}
	//改变文件大小为length
	_, err = m1.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	m1.length = length
	//进行系统调用实现共享内存
	m1.mmaphandle, err = windows.CreateFileMapping(
		windows.Handle(m1.file.Fd()),
		nil,
		prot,
		uint32(length>>32),
		uint32(length),
		nil)
	if err != nil {
		return nil, err
	}
	m1.addr, err = windows.MapViewOfFile(m1.mmaphandle, fileflag, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil, err
	}
	m = m1
	return m, nil
}

//关闭内存映射
func (m *Mmap) Close() (err error) {
	err = windows.UnmapViewOfFile(m.addr) //释放已映射空间
	if err != nil {
		return err
	}
	err = windows.CloseHandle(m.mmaphandle) //关闭文件映射对象
	if err != nil {
		return err
	}
	err = m.file.Close() //关闭文件描述符
	if err != nil {
		return err
	}
	return nil
}

//返回内存映射的空间首地址
func (m *Mmap) Addr() uintptr {
	return m.addr
}
