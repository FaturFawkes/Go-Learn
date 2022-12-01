package main

import "fmt"

func main(){

	var name = "Fatur"

	if name == "Fatur" {
		fmt.Println("Hai " + name)
	} else {
		fmt.Println("maaf tidak kenal")
	}

	// IF SHORT STATEMENT

	if length := len(name); length > 5{
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}