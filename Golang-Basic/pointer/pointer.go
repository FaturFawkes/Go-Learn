package main

import "fmt"

type Address struct{
	Kota, Provinsi, Negara string
}

type DOB struct {
	date string
}

func main() {
	// Address
	address1 := Address{"bogor", "jawa barat", "indonesia"}
	address2 := &address1
	address2.Kota = "jakarta"

	fmt.Println(address1)
	fmt.Println(address2)

	// DOB
	dob := DOB{
		date: "22 Oktober 2002",
	}
	// dob.date reference ke struct date
	// Jika data pada pointer berubah maka data yang reference ke pointer tersebut juga akan berubah
	date2 := &dob
	dob.date = "14 Juli 1999"
	fmt.Println(dob.date)
	fmt.Println(date2.date)
}