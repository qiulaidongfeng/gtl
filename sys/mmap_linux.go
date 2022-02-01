// mmap_linux
package sys

import (
	"os"
	"sync"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	READ  = unix.PROT_EXEC
	WRITE = unix.PROT_WRITE
	EXEC  = unix.PROT_EXEC
	RWX   = READ | WRITE | EXEC
)

type Mmap struct {
	mutex  sync.RWMutex
	file   *os.File
	length int
	memory []byte
}

func NewMmap(path string, length int) (m *Mmap, err error) {
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
	m.memory, err = unix.Mmap(int(m.file.Fd()), 0, length, RWX, unix.MAP_SHARED)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Mmap) Close() (err error) {
	m.mutex.Lock()         //加锁，保证并发安全
	defer m.mutex.Unlock() //解锁，保证并发安全
	err = unix.Munmap(m.memory)
	if err != nil {
		return err
	}
	err = m.file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *Mmap) Read(p []byte) (n int, err error) {
	pn := len(p) - 1
	m.mutex.RLock() //加锁，保证并发安全
	for n = 0; n < pn; n++ {
		p[n] = m.memory[n]
	}
	m.mutex.RUnlock() //解锁，保证并发安全
	return n, nil
}

func (m *Mmap) Write(p []byte) (n int, err error) {
	pn := len(p) - 1
	if pn < m.length {
		return 0, spancesmall
	}
	m.mutex.Lock() //加锁，保证并发安全
	for n = 0; n < pn; n++ {
		m.memory[n] = p[n]
	}
	m.mutex.Unlock() //解锁，保证并发安全
	return n, nil
}

func (m *Mmap) ReadAt(p []byte, off int64) (n int, err error) {
	pn := len(p) - 1
	m.mutex.Lock() //加锁，保证并发安全
	for n = int(off); n < pn; n++ {
		p[n] = m.memory[n]
	}
	m.mutex.Unlock() //解锁，保证并发安全
	n = m.length - n
	return n, nil
}

func (m *Mmap) WriteAt(p []byte, off int64) (n int, err error) {
	pn := len(p)
	if pn < m.length {
		return 0, spancesmall
	}
	m.mutex.Lock() //加锁，保证并发安全
	for n = int(off); n < pn; n++ {
		m.memory[n] = p[n]
	}
	n = m.length - n
	m.mutex.Unlock() //解锁，保证并发安全
	return n, nil
}

func (m *Mmap) Addr() uintptr {
	return uintptr(unsafe.Pointer(&m.memory[0]))
}

func (m *Mmap) AddrAt(off int) uintptr {
	return uintptr(unsafe.Pointer(&m.memory[off]))
}
