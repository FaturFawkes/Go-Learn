package main

import "fmt"

// MENGGUNAKAN FOR LOOP
func faktorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// MENGGUNAKAN RECURSIVE FUNCTION
func faktorialRecursive(val int) int {
	if val == 1 {
		return 1
	} else {
		return val * faktorialRecursive(val-1)
	}
}

func main() {
	fmt.Println(faktorialLoop(20))
	fmt.Println(faktorialRecursive(20))
}

