//go:build go1.16
// +build go1.16

package cachefs

import (
	"bytes"
	"io"
	"io/fs"
	"os"
	"sync"
	"time"
	_ "unsafe"
)

type HttpCacheFs struct {
	fd sync.Map //key=string value=*CacheFs
}

func NewHttpCacheFs() *HttpCacheFs {
	return &HttpCacheFs{}
}

func (fs *HttpCacheFs) Open(name string) (fs.File, error) {
	value, ok := fs.fd.Load(name)
	if ok {
		return value.(*CacheFs), nil
	}
	cache, err := NewCacheFs(name)
	if err != nil {
		return nil, err
	}
	return cache, nil
}

type CacheFs struct {
	fd      *os.File
	modtime time.Time
	buf     *bytes.Buffer
}

func NewCacheFs(name string) (*CacheFs, error) {
	fd, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	fdinfo, err := fd.Stat()
	if err != nil {
		return nil, err
	}
	modtime := fdinfo.ModTime()
	file, err := readall(fd)
	if err != nil {
		return nil, err
	}
	ret := &CacheFs{fd: fd, modtime: modtime, buf: bytes.NewBuffer(file)}
	return ret, nil
}

func (fs *CacheFs) Read(p []byte) (n int, err error) {
	fdinfo, err := fs.fd.Stat()
	if err != nil {
		return 0, err
	}
	nowtime := fdinfo.ModTime()
	if fs.modtime.Equal(nowtime) {
		return fs.buf.Read(p)
	}
	nowfs, err := NewCacheFs(fs.fd.Name())
	if err != nil {
		return 0, err
	}
	fs = nowfs
	return fs.Read(p)
}

func (fs *CacheFs) Close() error {
	err := fs.fd.Close()
	return err
}

func (fs *CacheFs) Stat() (fs.FileInfo, error) {
	return fs.fd.Stat()
}

//go:linkname readall io.ReadAll
func readall(r io.Reader) ([]byte, error)
