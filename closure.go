package main

import "fmt"

// value dalam variabel di dalam function mengubah value pada variable diatasnya jika namanya sama

func main() {
	counter := 0
	increment := func(){
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}