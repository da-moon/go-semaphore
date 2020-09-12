package semaphore_test

import (
	"strconv"
	"sync"
	"testing"

	semaphore "github.com/da-moon/go-semaphore"
)

func BenchmarkBinarySemaphore(b *testing.B) {
	lf := semaphore.NewBinarySemaphore()

	b.ReportAllocs()

	var cases = []int{1, 2, 4, 8, 16}
	for _, concurrency := range cases {
		b.Run(strconv.Itoa(concurrency), func(b *testing.B) {
			b.ResetTimer()
			b.SetParallelism(concurrency)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					lf.Wait()
					lf.Signal()
				}
			})
		})
	}

}

func BenchmarkMutex(b *testing.B) {
	m := sync.Mutex{}
	b.ReportAllocs()
	var cases = []int{1, 2, 4, 8, 16}
	for _, concurrency := range cases {
		b.Run(strconv.Itoa(concurrency), func(b *testing.B) {
			b.ResetTimer()
			b.SetParallelism(concurrency)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					m.Lock()
					m.Unlock()
				}
			})
		})
	}

}

// func BenchmarkCountingSemaphore(b *testing.B) {
// 	sem := newSlowSemaphore()
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	var cases = []int{1, 2, 4, 8, 16}
// 	for _, concurrency := range cases {
// 		b.Run(strconv.Itoa(concurrency), func(b *testing.B) {
// 			// b.SetParallelism(concurrency)
// 			b.SetParallelism(1)
// 			b.RunParallel(func(pb *testing.PB) {
// 				for pb.Next() {
// 					sem.Wait()
// 					sem.Signal()
// 				}
// 			})
// 		})
// 	}
// }
