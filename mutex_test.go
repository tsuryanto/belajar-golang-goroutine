package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {

	var mutex sync.Mutex

	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// Mutex untuk lock data
				/*
					digunakan untuk menandai bahwa semua operasi pada baris setelah kode tersebut adalah bersifat eksklusif.
					Hanya ada satu buah goroutine yang bisa melakukannya dalam satu waktu.
					Jika ada banyak goroutine yang eksekusinya bersamaan, harus antri.
				*/
				mutex.Lock()

				// Ini hanya bisa diakses oleh satu goroutine saja.
				x++

				// Mutex untuk Unlock data, jadi bisa diakses oleh goroutine yg lain
				mutex.Unlock()

			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Total x = ", x) // Hasilnya x = 100000 (benar)
}
