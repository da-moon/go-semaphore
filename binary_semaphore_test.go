package semaphore_test

import (
	"reflect"
	"sync"
	"testing"

	semaphore "github.com/da-moon/go-semaphore"
)

const size = 1000

func TestBinarySemaphore(t *testing.T) {
	lf := semaphore.NewBinarySemaphore()
	cnt := 0
	nums := make([]int, 0, size)
	wg := new(sync.WaitGroup)

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lf.Wait()
			nums = append(nums, cnt)
			cnt++
			lf.Signal()
		}(i)
	}

	wg.Wait()

	expected := make([]int, 0, size)
	for i := 0; i < size; i++ {
		expected = append(expected, i)
	}

	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("semphore is wrong. expected: %v, got: %v", expected, nums)
	}
}
