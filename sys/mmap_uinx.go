//go:build aix || darwin || freebsd || linux || netbsd || openbsd || solaris || dragonfly
// +build aix darwin freebsd linux netbsd openbsd solaris dragonfly

package sys

import (
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	//页面可读取
	READ = unix.PROT_READ
	//页面可写入
	WRITE = unix.PROT_WRITE
	//页面可执行
	EXEC = unix.PROT_EXEC
	RWX  = READ | WRITE | EXEC
)

//内存映射的结构体
type Mmap struct {
	file   *os.File
	length int
	memory []byte
}

//以读写模式打开文件，0777权限位，可读可写可执行
func NewMmap(path string, length int) (m *Mmap, err error) {
	NewMmapAll(path, os.O_RDWR|os.O_CREATE, 0777, length, RWX, unix.MAP_SHARED)
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
	m.memory, err = unix.Mmap(int(m.file.Fd()), 0, length, RWX, unix.MAP_SHARED)
	if err != nil {
		return nil, err
	}
	return m, nil*/
}

//以自定义模式与自定义权限位打开文件，自定义是否读写执行
func NewMmapAll(path string, osflag int, perm os.FileMode, length int, prot int, fileflag int) (m *Mmap, err error) {
	//打开文件
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
	m.memory, err = unix.Mmap(int(m.file.Fd()), 0, length, prot, fileflag)
	if err != nil {
		return nil, err
	}
	return m, nil
}

//关闭内存映射
func (m *Mmap) Close() (err error) {
	err = unix.Munmap(m.memory) //释放已映射空间
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
