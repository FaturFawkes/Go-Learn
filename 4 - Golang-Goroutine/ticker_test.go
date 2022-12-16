package belajargoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func ()  {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()
	
	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	
	for time := range channel {
		fmt.Println(time)
	}
}

func TestTickerSelect(t *testing.T) {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func ()  {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <- done :
			// ticker.Stop()
			return
		case time := <- ticker.C:
			fmt.Println(time)
		}
	}

}