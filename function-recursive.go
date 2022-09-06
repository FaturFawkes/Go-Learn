package main

import "fmt"

// MENGGUNAKAN FOR LOOP
func faktorialLoop(value int)int{
	result := 1
	for i := value; i >0 ; i-- {
		result *= i
	}
	return result
}

// MENGGUNAKAN RECURSIVE FUNCTION
func faktorialResursive(value int)int {
	if value == 1{
		return 1
	} else {
		return value * faktorialResursive(value-1)
	}
}

func main() {
	hasil := faktorialLoop(5)
	fmt.Println(hasil)

	fmt.Println(faktorialResursive(5))
}

