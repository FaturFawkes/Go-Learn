package main

import "fmt"

// Cara deklarasi struct
type Customer struct {
	Nama, Alamat string
	Umur int
}


func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "My name is", customer.Nama)
}


func main() {
	// Cara membuat struct 1
	var fatur Customer
	fatur.Nama = "Fatur"
	fatur.Alamat = "Jonggol"
	fatur.Umur = 19

	fatur.sayHello("rohman")

	// // Cara membuat struct 2
	// rohman := Customer{
	// Nama : "rohman",
	// Alamat : "Jonggol",
	// Umur : 19,
	// }

	// fmt.Println(fatur)
	// fmt.Println(rohman)
}