package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// Waiting 1 Goroutine until End

func printNumber(group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 100000; i++ {
		fmt.Println("Number ", i)
	}
}

func TestGoroutine(t *testing.T) {

	var group sync.WaitGroup
	group.Add(1)

	go printNumber(&group)

	group.Wait()
}
