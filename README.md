# go-semaphore

<p align="center">
  <a href="https://gitpod.io#https://github.com/da-moon/go-semaphore">
    <img src="https://img.shields.io/badge/open%20in-gitpod-blue?logo=gitpod" alt="Open In GitPod">
  </a>
  <img src="https://img.shields.io/github/languages/code-size/da-moon/go-semaphore" alt="GitHub code size in bytes">
  <img src="https://img.shields.io/github/commit-activity/w/da-moon/go-semaphore" alt="GitHub commit activity">
  <img src="https://img.shields.io/github/last-commit/da-moon/go-semaphore/master" alt="GitHub last commit">
</p>

spinlock based semaphore in go
this package is still experimental. use at your own risk.

## benchmarks

- benchmark results (go version `go1.15 linux/amd64`)

```
Running tool: /home/gitpod/go/bin/go test -benchmem -run=^$ github.com/da-moon/go-semaphore -bench . -v
goos: linux
goarch: amd64
pkg: github.com/da-moon/go-semaphore
BenchmarkBinarySemaphore
BenchmarkBinarySemaphore/1
BenchmarkBinarySemaphore/1-16              	32325340	        37.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinarySemaphore/2
BenchmarkBinarySemaphore/2-16              	26361247	        39.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinarySemaphore/4
BenchmarkBinarySemaphore/4-16              	30186558	        37.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinarySemaphore/8
BenchmarkBinarySemaphore/8-16              	33298560	        38.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinarySemaphore/16
BenchmarkBinarySemaphore/16-16             	25791638	        41.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutex
BenchmarkMutex/1
BenchmarkMutex/1-16                        	15966688	        78.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutex/2
BenchmarkMutex/2-16                        	13372855	        83.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutex/4
BenchmarkMutex/4-16                        	19589612	        85.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutex/8
BenchmarkMutex/8-16                        	12234087	        96.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutex/16
BenchmarkMutex/16-16                       	11732821	        89.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlowSemaphorePermitpool200
BenchmarkSlowSemaphorePermitpool200-16     	     100	  10524237 ns/op	  886794 B/op	  110634 allocs/op
BenchmarkSemaphore200
BenchmarkSemaphore200-16                   	     272	   4087844 ns/op	    1184 B/op	       3 allocs/op
BenchmarkSlowSemaphorePermitpool500
BenchmarkSlowSemaphorePermitpool500-16     	      16	  75324527 ns/op	 5827348 B/op	  726940 allocs/op
BenchmarkSemaphore500
BenchmarkSemaphore500-16                   	      43	  25890357 ns/op	    1434 B/op	       3 allocs/op
BenchmarkSlowSemaphorePermitpool1000
BenchmarkSlowSemaphorePermitpool1000-16    	       3	 339437950 ns/op	24780800 B/op	 3090996 allocs/op
BenchmarkSemaphore1000
BenchmarkSemaphore1000-16                  	      10	 106566390 ns/op	    2720 B/op	       7 allocs/op
BenchmarkSlowSemaphorePermitpool3000
BenchmarkSlowSemaphorePermitpool3000-16    	       1	3106959920 ns/op	214735584 B/op	26684397 allocs/op
BenchmarkSemaphore3000
BenchmarkSemaphore3000-16                  	       2	 875797314 ns/op	    1184 B/op	       3 allocs/op
```

