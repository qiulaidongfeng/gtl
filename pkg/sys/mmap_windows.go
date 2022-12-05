package sys

import (
	"os"
	"syscall"
)

const (
	//页面只读
	PAGE_READ = syscall.PAGE_READONLY
	//页面读写
	PAGE_RW = syscall.PAGE_READWRITE
	//页面可读，可写，可执行
	PAGE_RE = syscall.PAGE_EXECUTE_READWRITE
	//写复制
	PAGE_WRITECOPY = syscall.PAGE_EXECUTE_WRITECOPY
)

const (
	//写复制
	FILE_MAP_COPY = syscall.FILE_MAP_COPY
	//读写
	FILE_MAP_WRITE = syscall.FILE_MAP_WRITE
	//只读
	FILE_MAP_READ = syscall.FILE_MAP_READ
)

var (
	//系统内存页的尺寸
	Pagesize = uint(os.Getpagesize())
)

//内存映射的结构体
type Mmap struct {
	file       *os.File
	mmaphandle Handle
	length     uint
	addr       uintptr
}

//以读写模式打开文件，0777权限位，可读可写
func NewMmap(path string, length uint) (m *Mmap, err error) {
	m, err = NewMmapAll(path, os.O_RDWR|os.O_CREATE, 0777, length, PAGE_RW, FILE_MAP_WRITE)
	return
}

//以自定义模式与自定义权限位打开文件，自定义是否读写执行
func NewMmapAll(path string, osflag int, perm os.FileMode, length uint, prot uint32, fileflag uint32) (m *Mmap, err error) {
	var m1 *Mmap = new(Mmap)
	//打开文件
	m1.file, err = os.OpenFile(path, osflag, perm)
	if err != nil {
		return nil, err
	}
	//移动fd到开头
	_, err = m1.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	m1.length = length
	//进行系统调用实现共享内存
	m1.mmaphandle, err = syscall.CreateFileMapping(
		Handle(m1.file.Fd()),
		nil,
		prot,
		uint32(length>>32),
		uint32(length),
		nil)
	if err != nil {
		return nil, err
	}
	m1.addr, err = syscall.MapViewOfFile(m1.mmaphandle, fileflag, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	m = m1
	return m, nil
}

//关闭内存映射
func (m *Mmap) Close() (err error) {
	err = syscall.UnmapViewOfFile(m.addr) //释放已映射空间
	if err != nil {
		return err
	}
	err = syscall.CloseHandle(m.mmaphandle) //关闭文件映射对象
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
