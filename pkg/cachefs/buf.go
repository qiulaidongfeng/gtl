package cachefs

import (
	"io"
)

var _ io.Reader = (*Buf)(nil)

// 循环缓存
// 当到末尾是不会返回io.EoF,会下次从头读取
type Buf struct {
	i      int
	buflen int
	buf    []byte
}

// 创建循环缓存
func NewBuf(buf []byte) *Buf {
	ret := &Buf{buf: buf}
	ret.buflen = len(buf)
	return ret
}

// 实现io.Reader接口,读取缓存内容
// 当到末尾是不会返回io.EoF,会下次从头读取
func (buf *Buf) Read(p []byte) (n int, err error) {
	if buf.Empty() {
		if len(p) == 0 {
			return 0, nil
		}
		buf.i = 0
	}
	n = copy(p, buf.buf[buf.i:])
	buf.i += n
	return n, nil
}

// 实现io.Seeker接口
func (buf *Buf) Seek(offset int64, whence int) (int64, error) {
	off := int(offset)
	switch whence {
	case io.SeekStart:
		buf.i = off
	case io.SeekCurrent:
		buf.i += off
	case io.SeekEnd:
		buf.i = buf.buflen - off
	}
	return int64(buf.i), nil
}

// 判断缓存是否到末尾
func (buf *Buf) Empty() bool {
	return buf.i >= buf.buflen
}
