//go:build !windows
// +build !windows

package sys

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	//页面可读取
	READ = syscall.PROT_READ
	//页面可写入
	WRITE = syscall.PROT_WRITE
	RW    = READ | WRITE
)

const (
	//写入数据复制回文件内，允许其他映射该文件的进程共享
	SHARED = syscall.MAP_SHARED
	//写入时复制,写入操作会产生映射文件的复制，对此的任何修改都不会写入文件
	PRIVATE = syscall.MAP_PRIVATE
)

var (
	Pagesize = os.Getpagesize()
)

//内存映射的结构体
type Mmap struct {
	file   *os.File
	length uint
	memory []byte
}

//以读写模式打开文件，0777权限位，可读可写可执行
func NewMmap(path string, length uint) (m *Mmap, err error) {
	m, err = NewMmapAll(path, os.O_RDWR|os.O_CREATE, 0777, length, RW, SHARED)
	return m, err
}

//以自定义模式与自定义权限位打开文件，自定义是否读写执行
func NewMmapAll(path string, osflag int, perm os.FileMode, length uint, prot int, fileflag int) (m *Mmap, err error) {
	m = new(Mmap)
	//打开文件
	m.file, err = os.OpenFile(path, osflag, perm)
	if err != nil {
		return nil, err
	}
	//改变文件大小为length
	_, err = m.file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	m.length = length
	//进行系统调用实现共享内存
	m.memory, err = syscall.Mmap(int(m.file.Fd()), 0, int(length), prot, fileflag)
	if err != nil {
		return nil, err
	}
	return m, nil
}

//关闭内存映射
func (m *Mmap) Close() (err error) {
	err = syscall.Munmap(m.memory) //释放已映射空间
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
	return uintptr(unsafe.Pointer(&m.memory[0]))
}
