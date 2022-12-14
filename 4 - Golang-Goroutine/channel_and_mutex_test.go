package belajargoroutine

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)



func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fatur"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second) 
	channel <- "Bnox"
}

// Channel as parameter
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	
	go OnlyIn(channel)
	go OnlyOut(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// Channel in/out only

// Only in
func OnlyIn(channel chan<- string) {
	channel <- "Fatur"
}

// Only out
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

// Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func ()  {
		channel <- "Nur"
		channel <- "Fatur"
		channel <- "Rohman"
	}()
	
	go func ()  {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
}

// Range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		for i := 0; i < 20; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data ", data)
	}
	fmt.Println("Selesai")
}

// Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	var counter int
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}

// Race Condition
func TestRaceCondition(t *testing.T) {
	var x int
	var mutex sync.RWMutex

	for i := 1; i <= 1000; i++ {
		go func ()  {
			for j := 1; j <= 100; j++ {
				mutex.RLock()
				x += 1
				mutex.RUnlock()
				fmt.Println(x)
			}
			}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}