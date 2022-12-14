package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	group := &sync.WaitGroup{}
	once := &sync.Once{}



	for i := 0; i < 100; i++ {
		go func ()  {
			group.Add(1)
				once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
	fmt.Println("Selesai")

}