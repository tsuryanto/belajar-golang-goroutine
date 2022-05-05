package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func onlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(onlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter :", counter)

}
