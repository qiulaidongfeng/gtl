package gsync

import (
	"sync"
)

//读写锁接口
type Rwlock interface {
	sync.Locker
	RLock()
	RUnlock()
}
