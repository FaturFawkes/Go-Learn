package main

import "fmt"

func main(){
	// CARA 1 BUAT ARRAY
	var nama[3]string

	nama[0] = "Nur"
	nama[1] = "Fatur"
	nama[2] = "Rohman"

	fmt.Println(nama[1])

	// CARA 2 BUAT ARRAY
	var kelas = [6]int{
		1, 2, 3, 4, 5, 6,
	}

	fmt.Println(kelas)
}