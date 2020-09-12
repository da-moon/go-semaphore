package semaphore

// Semaphore - main semaphore interface
type Semaphore interface {
	Signal()
	Wait()
}
