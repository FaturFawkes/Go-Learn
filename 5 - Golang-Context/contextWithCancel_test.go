package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	fmt.Println(runtime.NumGoroutine())

	go func ()  {
		defer close(destination)	
		counter := 1
		for {
			select {
			case <- ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}

		}
	}()

	fmt.Println(runtime.NumGoroutine())
	return destination
}

func TestCtxWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)


	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter :",n)
		if n == 10 {
			break
		}
	}
	cancel() // MENGIRIM SINYAL CANCEL KE CONTEXT

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}