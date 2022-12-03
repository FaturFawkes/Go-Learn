package main

import "fmt"

func random() interface{} {
	return false
}

func catch(){
	if r := recover(); r != nil{
		fmt.Println("app error")
	} 
}

func main() {
	defer catch()
	res := random()

	// MENGECEK TIPE DATA DARI INTERFACE DENGAN SWITCH
	switch res.(type) {
	case string:
		fmt.Println("Ini adalah string", res)
	case bool:
		fmt.Println("Ini adalah boolean", res)
	default:
		panic("error")
	} 

	// MENGECEK LANGSUNG TIPE DATA
	// Jika sudah yakin return tipe datanya apa
	tipe := res.(bool)
	fmt.Println(tipe, res)
}