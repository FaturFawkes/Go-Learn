package main

import "fmt"

func main(){
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice = months[4:7]
	
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// months[5] = "Diubah"
	// fmt.Println(slice)

	// slice[0] = "Mei Lagi"
	// fmt.Println(slice)

	var slice1 = months[11:]
	fmt.Println(slice1)

	var slice2 = append(slice1, "Fatur")
	fmt.Println(slice2)
	slice2[1] = "Bukan Desember"
	fmt.Println(slice2)

	fmt.Println(slice1)
	fmt.Println(months)

	var newSlice = make([]string, 2, 5)

	newSlice[0] = "Fatur"
	newSlice[1] = "Rohman"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}