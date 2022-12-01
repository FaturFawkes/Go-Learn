package main

import "fmt"

func main(){
	var name = "Fatur"
	// name = "Joko"

	switch name{
	case "Fatur" :
		fmt.Println("Halo " + name)
	case "Joko" :
		fmt.Println("Halo " + name)
	default :
		fmt.Println("Hi boleh kenalan?")	
	}

	// SWITCH SHORT STATEMENT

	// switch length := len(name); length > 5{
	// case true :
	// 	fmt.Println("Nama terlalu panjang")
	// case false :
	// 	fmt.Println("Nama sudah benar")
	// }

	// SWITCH TANPA DEKLARASI KONDISI
	
	name = "fatur rohman ganteng idaman"
	length := len(name)
	
	switch{
	case length < 5 :
		fmt.Println("Length < 5")
	case length <= 10 :
		fmt.Println("Length <= 10")
	default :
		fmt.Println("Length > 15")
}

}