package semaphore_test

import (
	"sync"
	"testing"

	semaphore "github.com/da-moon/go-semaphore"
)

// COMMENT :
// almost identical to hashicorp's lib as far as performance goes
func permitpoolRun(n int) int64 {
	var a int64 = 0
	sem := newSlowSemaphore()
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				sem.Signal()
				sem.Wait()

			}
			wg.Done()
		}()
	}
	wg.Wait()
	return a
}
func semaphoreRun(n int) int64 {
	var a int64 = 0
	sem := semaphore.NewCountingSemaphore()
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				sem.Signal()
				sem.Wait()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return a
}
func semaphoreBench(c int, b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		semaphoreRun(c)
	}
}
func permitpoolBench(c int, b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		permitpoolRun(c)
	}
}

func BenchmarkSlowSemaphorePermitpool200(b *testing.B) {
	permitpoolBench(200, b)

}

func BenchmarkSemaphore200(b *testing.B) {
	semaphoreBench(200, b)

}

func BenchmarkSlowSemaphorePermitpool500(b *testing.B) {
	permitpoolBench(500, b)
}
func BenchmarkSemaphore500(b *testing.B) {
	semaphoreBench(500, b)
}

func BenchmarkSlowSemaphorePermitpool1000(b *testing.B) {
	permitpoolBench(1000, b)
}

func BenchmarkSemaphore1000(b *testing.B) {
	semaphoreBench(1000, b)
}

func BenchmarkSlowSemaphorePermitpool3000(b *testing.B) {
	permitpoolBench(3000, b)
}

func BenchmarkSemaphore3000(b *testing.B) {
	semaphoreBench(3000, b)
}
