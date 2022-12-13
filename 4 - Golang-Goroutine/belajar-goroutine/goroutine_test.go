package belajargoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World!")
}

func DisplayNumber(num int) {
	fmt.Println("Display - ", num)
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Halo kambing")

	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
