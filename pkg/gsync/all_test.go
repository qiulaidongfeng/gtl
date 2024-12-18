package gsync

import "testing"

func TestRWMutex(t *testing.T) {
	lock := LMrwmutex{}
	lock.RLock()
	lock.RUnlock()
	lock.Lock()
}
