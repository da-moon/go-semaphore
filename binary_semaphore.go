package semaphore

import (
	"runtime"
	"sync/atomic"
)

type state int32

const (
	idle      state = 0
	modifying state = 1
)

// countingSemaphore - main
type binarySemaphore struct {
	state int32
}

// NewBinarySemaphore returns a new binary semaphore
func NewBinarySemaphore() Semaphore {
	return &binarySemaphore{
		state: 0,
	}
}

// Wait waits until the shared counter is available. Update the sharing counter if available
func (g *binarySemaphore) Wait() {
	for !atomic.CompareAndSwapInt32(&g.state, int32(idle), int32(modifying)) {
		runtime.Gosched()
	}
}

// Signal signals termination of use of shared counter
func (g *binarySemaphore) Signal() {
	for !atomic.CompareAndSwapInt32(&g.state, int32(modifying), int32(idle)) {
		runtime.Gosched()
	}
}
