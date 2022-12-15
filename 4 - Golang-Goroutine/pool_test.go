package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() any {
			return "New"
		},
	}
	var wait sync.WaitGroup

	pool.Put("Nur")
	pool.Put("Fatur")
	pool.Put("Rohman")

	for i := 0; i < 10; i++ {
		go func ()  {
			wait.Add(1)
			data := pool.Get()
			fmt.Println(data)
			wait.Done()
			pool.Put(data)
		}()
	}
	wait.Wait()
	fmt.Println("Selesai")
}