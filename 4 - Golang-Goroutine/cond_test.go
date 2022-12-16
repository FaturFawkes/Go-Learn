package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock() 
	cond.Wait()
	fmt.Println("Done", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}
	
	// Signal akan mengizinkan cond.Wait untuk jalan satu satu dan menunggu
	go func ()  {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// Broadcast akan mengizinakn cond.Wait untuk jalan semua tanpa harus menunggu laig
	// go func ()  {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()
	group.Wait()
}