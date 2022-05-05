package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func addToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go addToMap(data, i, &group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
