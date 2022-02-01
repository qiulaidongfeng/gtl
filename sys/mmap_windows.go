// mmap_windows
package sys

import (
	"os"

	"golang.org/x/sys/windows"
)

const (
	PAGE_READ = windows.PAGE_READONLY
	PAGE_RW   = windows.PAGE_READWRITE
	PAGE_EXEC = windows.PAGE_EXECUTE
	PAGE_RWX  = windows.PAGE_EXECUTE_READWRITE
)

const (
	FILE_MAP_COPY    = windows.FILE_MAP_COPY
	FILE_MAP_WRITE   = windows.FILE_MAP_WRITE
	FILE_MAP_READ    = windows.FILE_MAP_READ
	FILE_MAP_EXECUTE = windows.FILE_MAP_EXECUTE
	FILE_MAP_RWX     = FILE_MAP_READ | FILE_MAP_WRITE | FILE_MAP_EXECUTE
)

var (
	pagesize = uint(os.Getpagesize())
)

type Mmap struct {
	mutex  sync.RWMutex
	file   *os.File
	length uint
	addr   uintptr
}

func NewMmap(path string, length uint) (m *Mmap, err error) {
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
	handle := windows.Handle(0)
	handle, err = windows.CreateFileMapping(
		windows.Handle(m.file.Fd()),
		nil,
		RWX,
		uint32(length>>32),
		uint32(length),
		nil)
	if err != nil {
		return nil, err
	}
	//使文件大小等于内存页倍数
	pagelen := length%uint(pagesize) + 1
	size := pagelen * pagesize
	m.addr, err = windows.MapViewOfFile(handle, FILE_MAP_RWX, 0, 0, uintptr(size))
	if err != nil {
		return nil, err
	}
	return m, nil
}
