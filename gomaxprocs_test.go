package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {

	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	fmt.Println("Total CPU", runtime.NumCPU())

	// Jika ingin mengganti jumlah thread
	// runtime.GOMAXPROCS(20)

	fmt.Println("Total Thread", runtime.GOMAXPROCS(-1))
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	group.Wait()
}
