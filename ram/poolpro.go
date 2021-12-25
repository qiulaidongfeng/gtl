// poolpro
package ram

import (
	"errors"
	"sync"
	"sync/atomic"
)

type poolpro struct {
	pool  sync.Pool
	len   *uint64
	mutex sync.Mutex
}

func Newpoolpro() poolpro {
	p := poolpro{
		pool: *(new((sync.Pool))),
		len:  new(uint64),
	}
	return p
}

func (p poolpro) Len() uint64 {
	return *(p.len)
}

func (p *poolpro) Put(x interface{}) {
	*(p.len) += 1
	p.pool.Put(x)
	return
}

func (p *poolpro) Get() (interface{}, error) {
	if *(p.len) == 0 {
		return nil, errors.New("poolpro,Empty")
	}
	*(p.len) -= 1
	return p.pool.Get(), nil
}

func (p poolpro) Tslen() uint64 {
	return atomic.LoadUint64(p.len)
}

func (p *poolpro) Tsput(x interface{}) {
	_ = atomic.AddUint64(p.len, 1)
	p.mutex.Lock()
	p.pool.Put(x)
	p.mutex.Unlock()
	return
}

func (p *poolpro) Tsget() (interface{}, error) {
	if atomic.LoadUint64(p.len) == 0 {
		return nil, errors.New("poolpro,Empty")
	}
	_ = atomic.AddUint64(p.len, ^uint64(0))
	p.mutex.Lock()
	object := p.pool.Get()
	p.mutex.Unlock()
	return object, nil
}
