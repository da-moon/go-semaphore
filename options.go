package semaphore

// Option - options setter method
type Option func(*countingSemaphore)

// WithPermits - sets number of permits in the permit pool
func WithPermits(arg int) Option {
	return func(e *countingSemaphore) {
		e.permits = arg
	}
}
