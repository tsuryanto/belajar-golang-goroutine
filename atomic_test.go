package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {

	var group sync.WaitGroup

	var x int64 = 0
	for i := 0; i < 1000; i++ {

		group.Add(1)

		go func() {
			defer group.Done()

			for j := 0; j < 100; j++ {
				// x = x + 1
				/*
					Atomic digunakan agar tidak terjadi Race Condition.
					Jika data berupa primitive variable bisa menggunakan atomic,
					namun jika data struct / data yg cukup kompleks bisa menggunakan mutex
				*/
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Total x = ", x)
}
