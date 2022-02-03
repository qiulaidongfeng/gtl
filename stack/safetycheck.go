// safetycheck
package stack

import (
	"sync/atomic"
)

const (
	safeOk int8 = iota
	notEnoughSpace
)

func (s *GLMstack) pushsafetycheck(size uint64) (safe int8) {
	if s.size+size >= s.scap {
		safe = notEnoughSpace
	}
	return safeOk
}

func (s *GLMstack) tspushsafetycheck(size uint64) (sizeold uint64, safe int8, nsize uint64) {
	nsize = atomic.AddUint64(&s.size, size)
	if nsize >= s.scap {
		safe = notEnoughSpace
	}
	sizeold = nsize - size
	return
}
