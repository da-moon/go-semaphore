package semaphore

// countingSemaphore - main
type countingSemaphore struct {
	sem     chan int
	permits int
}

// New -
func NewCountingSemaphore(opts ...Option) Semaphore {
	result := &countingSemaphore{}
	for _, opt := range opts {
		opt(result)
	}
	// Default number of parallel operations
	if result.permits <= 1 {
		result.permits = 128
	}
	result.sem = make(chan int, result.permits)
	return result
}

// Signal - returns when a countingSemaphore has been Signald
func (s *countingSemaphore) Signal() {
	s.sem <- 1
}

// Wait - Waits a countingSemaphore
func (s *countingSemaphore) Wait() {
	<-s.sem
}
