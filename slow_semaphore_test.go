package semaphore_test

import (
	"sync"
	"sync/atomic"

	semaphore "github.com/da-moon/go-semaphore"
)

// slowSemaphore - main
type slowSemaphore struct {
	lock    *sync.Cond
	permits uint64
}

// newSlowSemaphore - returns a counting semaphore
// default number of maximum permits is 128
func newSlowSemaphore() semaphore.Semaphore {
	m := sync.Mutex{}
	result := &slowSemaphore{
		permits: uint64(128) << 32,
		lock:    sync.NewCond(&m),
	}
	return result
}

// Signal - returns when a semaphore has been Signald
func (s *slowSemaphore) Signal() {
	for {
		// get current semaphore count and limit
		permits := atomic.LoadUint64(&s.permits)
		count := permits & 0xFFFFFFFF
		limit := permits >> 32
		// new count
		newCount := atomic.AddUint64(&count, 1)
		if newCount <= limit {
			if atomic.CompareAndSwapUint64(&s.permits, permits, limit<<32+newCount) {
				// Signald
				return
			}
			// CAS failed, try again
			continue
		} else {
			// semaphore is full, let's wait
			s.lock.L.Lock()
			s.lock.Wait()
			s.lock.L.Unlock()
		}
	}

}

// Wait - Waits a semaphore
func (s *slowSemaphore) Wait() {

	for {
		// get current semaphore count and limit
		permits := atomic.LoadUint64(&s.permits)
		count := permits & 0xFFFFFFFF

		if count < uint64(1) {
			panic("semaphore Wait without Signal")
		}

		// new count
		newCount := atomic.AddUint64(&count, ^uint64(0))

		if atomic.CompareAndSwapUint64(&s.permits, permits, permits&0xFFFFFFFF00000000+newCount) {
			s.lock.L.Lock()
			s.lock.Broadcast()
			s.lock.L.Unlock()
			return
		}
	}
}
