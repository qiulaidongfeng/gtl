// gsync
package gsync

import (
	"sync"
)

type Rwlock interface {
	sync.Locker
	RLock()
	RUnlock()
}
