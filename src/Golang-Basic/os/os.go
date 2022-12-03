package main

import (
	"fmt"
	"os"
)

func main() {
	// ARGUMENT 
	// Untuk mengambil arugment ketika run file golang 
	argument := os.Args
	// fmt.Println(argument)
	fmt.Println("ARGUMENT : ", argument[1])

	// HOSTNAME
	// Untuk mengambil nama host/nama PC 
	host, _ := os.Hostname()
	fmt.Println("HOSTNAME : ", host)

}