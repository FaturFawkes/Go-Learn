package main 

import "fmt"

func main(){

	// Operator Boolean

	var nilaiA = 90
	var nilaiB = 80

	// var hasil1 bool 
	var hasil2 bool
	var hasil1 bool
	var hasil3 bool

	hasil1 = nilaiA > nilaiB
	hasil2 = nilaiA == nilaiB
	hasil3 = hasil1 && hasil2

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
	
}