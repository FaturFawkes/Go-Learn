package belajargoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64
	group := &sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}
	group.Wait()
	fmt.Println("Counter = ", x)
}
