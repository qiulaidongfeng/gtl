// safetycheck
package stack

import (
	"sync/atomic"
)

const (
	safeOk int8 = iota
	notEnoughSpace
	shortageStackContent
<<<<<<< HEAD
	ncapSmall
=======
>>>>>>> c91b074acab5ede5c2be5087800827f414808c76
)

func (s *GLMstack) pushsafetycheck(size uint64) (safe int8) {
	if s.size+size >= s.scap {
		safe = notEnoughSpace
	}
	return
}

func (s *GLMstack) tspushsafetycheck(size uint64) (sizeold uint64, safe int8, nsize uint64) {
	nsize = atomic.AddUint64(&s.size, size)
	if nsize >= s.scap {
		safe = notEnoughSpace
	}
	sizeold = nsize - size
	return
}

func (s *GLMstack) popsafetycheck(size uint64) (safe int8) {
	if size > s.size {
		safe = shortageStackContent
	}
	return
}
<<<<<<< HEAD

func (s *GLMstack) addcapsafetycheck(ncap uint64) (safe int8) {
	if ncap <= s.scap {
		safe = ncapSmall
	}
	return
}

func (s *GLMstack) subcapsafetycheck(ncap uint64) (safe int8) {
	if ncap >= s.scap {
		safe = ncapSmall
	}
	return
}
=======
>>>>>>> c91b074acab5ede5c2be5087800827f414808c76
