package main

import "fmt"

func sumAll(hitung ...int) int {
	total := 0

	for _, number := range hitung{
		total += number
	}

	return total
}

func main() {
	result := sumAll(10, 10, 10, 10, 10)
	fmt.Println(result)

	slice := []int{10, 10, 10, 10, 10, 15}
	fmt.Println(sumAll(slice...))
}