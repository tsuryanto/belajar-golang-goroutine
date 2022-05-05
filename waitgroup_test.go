package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// var counter int = 0

func RunAsynchronous(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1)

	// counter++
	fmt.Println("Hello", i)

	// time.Sleep(1 * time.Second)
}

// func RunAsynchronousNoMutex(group *sync.WaitGroup) {
// 	defer group.Done()

// 	group.Add(1)

// 	counter++
// 	fmt.Println("Hello", counter)

// 	// time.Sleep(1 * time.Second)
// }

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	// gunakan mutex agar counter di proses bergantian 1 - 100
	// var mutex sync.Mutex

	for i := 0; i < 100000; i++ {
		go RunAsynchronous(group, i)
		// go RunAsynchronousNoMutex(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}

func TestGoLoop(t *testing.T) {
	for i := 0; i < 100000; i++ {
		fmt.Println("Hello", i)
	}
}
