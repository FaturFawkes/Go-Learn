package main

import "fmt"

type Address struct{
	Kota, Provinsi, Negara string
}

func main() {
	address1 := Address{"bogor", "jawa barat", "indonesia"}
	address2 := &address1
	address2.Kota = "jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}