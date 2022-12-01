package main

import "fmt"

func main(){

	// Konversi INTEGER
	var nilai1 int32 = 1000
	var nilai2 int64 = int64(nilai1)
	var nilai3 int8 = int8(nilai1)

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai3)


	// Konversi STRING
	var nama = "Nur Fatchurohman"
	var karakter = nama[0]
	var konversi = string(karakter)

	fmt.Println(konversi)

}



