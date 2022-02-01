// mmap_windows
package sys

import (
	"golang.org/x/sys/windows"
)

type Mmap struct {
	mutex  sync.RWMutex
	file   *os.File
	length int
	addr   uintptr
}
