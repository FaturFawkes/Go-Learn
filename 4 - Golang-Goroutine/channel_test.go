package belajargoroutine

import (
	"fmt"
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