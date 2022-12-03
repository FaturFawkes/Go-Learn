package main

import "fmt"

func main(){


	// MIRIP SEPERTI ARRAY ASOSIATIF
	
	var person = map[string]string{
		"name" : "Fatur",
		"address" : "Jonggol",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	fmt.Println(len(person))

	// Membuat map tanpa isi

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "Rohman"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}