//go:build go1.16
// +build go1.16

//cachefs实现了缓存文件系统
//
//缓存文件系统用于在http.FileSystem默认实现需要优化
//缓存文件系统在Open和Read时如果没有修改，不会进行系统调用，而是使用缓存
package cachefs

import (
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"
	_ "unsafe"
)

//Http缓存文件系统
type HttpCacheFs struct {
	path string
	fd   map[string]*CacheFs
}

//创建NewHttpCacheFs
//path 是相对的路径 相当于http.Dir
func NewHttpCacheFs(path string) *HttpCacheFs {
	return &HttpCacheFs{path: path, fd: make(map[string]*CacheFs)}
}

//Open实现http.FileSystem接口
func (fs *HttpCacheFs) Open(name string) (http.File, error) {
	fdname := filepath.Join(fs.path, name)
	value, ok := fs.fd[fdname]
	if ok {
		return value, nil
	} else {
		cache, err := NewCacheFs(fdname)
		if err != nil {
			return nil, err
		}
		fs.fd[fdname] = cache
		return cache, nil
	}
	return nil, nil
}

//缓存文件系统
//在Open和Read时如果没有修改，不会进行系统调用，而是使用缓存
type CacheFs struct {
	fd      *os.File
	modtime time.Time //缓存创建时修改时间
	buf     *Buf
}

//创建缓存文件系统
func NewCacheFs(name string) (fs *CacheFs, err error) {
	fd, err := os.Open(name) //打开
	if err != nil {
		return nil, err
	}
	fdinfo, err := fd.Stat()
	if err != nil {
		return nil, err
	}
	modtime := fdinfo.ModTime() //获取修改时间
	file, err := readall(fd)    //读取全部内容
	if err != nil {
		return nil, err
	}
	ret := &CacheFs{fd: fd, modtime: modtime, buf: NewBuf(file)} //创建缓存文件系统，文件全部内容被缓存
	return ret, nil
}

//实现io.Reader接口，如果没有修改，将返回缓存内容
func (fs *CacheFs) Read(p []byte) (n int, err error) {
	fdinfo, err := fs.fd.Stat()
	if err != nil {
		return 0, err
	}
	nowtime := fdinfo.ModTime()    //获取文件现在修改时间
	if fs.modtime.Equal(nowtime) { //通过比较缓存时修改时间，判断是否修改，没有直接从缓存读取
		return fs.buf.Read(p)
	}
	//有修改，重现缓存再读
	nowfs, err := NewCacheFs(fs.fd.Name())
	if err != nil {
		return 0, err
	}
	fs = nowfs
	return fs.Read(p)
}

//实现io.Seeker接口，如果没有修改，将移动缓存内容偏移量
func (fs *CacheFs) Seek(offset int64, whence int) (int64, error) {
	fdinfo, err := fs.fd.Stat()
	if err != nil {
		return 0, err
	}
	nowtime := fdinfo.ModTime()    //获取文件现在修改时间
	if fs.modtime.Equal(nowtime) { //通过比较缓存时修改时间，判断是否修改，没有直接移动缓存内容偏移量
		return fs.buf.Seek(offset, whence)
	}
	//有修改，重现缓存再移动缓存内容偏移量
	nowfs, err := NewCacheFs(fs.fd.Name())
	if err != nil {
		return 0, err
	}
	fs = nowfs
	return fs.Seek(offset, whence)
}

//返回目录信息
//目前返回nil,nil
func (fs *CacheFs) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}

//关闭
//为了能配合http.FileServer，永远返回nil
func (fs *CacheFs) Close() error {
	return nil
}

//返回文件信息
func (fs *CacheFs) Stat() (fs.FileInfo, error) {
	return fs.fd.Stat()
}

//go:linkname readall io.ReadAll
func readall(r io.Reader) ([]byte, error)
