package goroutine

import (
	"fmt"
	"strconv"
	"testing"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	sayHelloTo := func(who string) {
		// data dimasukkan ke channel
		channel <- "Hello " + who
	}

	go sayHelloTo("John")
	go sayHelloTo("Taufiq")

	// Mengambil data dari channel
	var channel1 = <-channel
	fmt.Println(channel1)

	// Mengambil data dari channel
	var channel2 = <-channel
	fmt.Println(channel2)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke " + strconv.Itoa(i)
		}

		// wajib di close untuk menghentikan perulangan
		// Jika tidak akan terjadi deadlock saat iterasi selesai
		close(channel)
	}()

	// for range disini akan memantai data masuk ke channel
	for data := range channel {
		fmt.Println("Menerima Data ", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}
