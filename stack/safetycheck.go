// safetycheck
package stack

const (
	safeOk         int8 = itoa
	notEnoughSpace      = 1
)

func (s *GLMstack) pushsafetycheck(size uint64) (safe int8) {
	if s.size+size >= s.scap {
		safe = notEnoughSpace
	}
	return safeOk
}

func (s *GLMstack) tspushsafetycheck(size uint64) (sizeold uint64) {
	nsize := atomic.AddUint64(&s.size, size)
	if nsize >= s.scap {
		s.mutex.RUnlock()
		s.scap = s.tsaddcap(nsize)
		s.mutex.RLock()
	}
	sizeold = nsize - size
	return
}
