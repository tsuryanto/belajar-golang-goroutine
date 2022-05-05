package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
		}()
	}

	time.Sleep(3 * time.Second)

	// harusnya x = 100000
	// namun tatalnya tidak sampai segitu (misal hanya 93432) karena terjadi Race Condition
	// Kondisi dimana goroutine secara bersamaan mengubah x
	/*
		Misal :
			Goroutine 1 -> x=1000 + 1
			Goroutine 2 -> x=1000 + 1
			Goroutine 3 -> x=1000 + 1
		Jadi ke 3 goroutine mengakses x secara bersamaan lalu mengubah menjadi value yang sama

		Cara mengatasinya :
		1. Menggunakan sync.Mutex
		2. Menggunakan Atomic untuk primitive variable
	*/
	fmt.Println("Total x = ", x)
}
