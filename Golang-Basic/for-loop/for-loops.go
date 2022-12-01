package main

import "fmt"

func main(){
	// var counter = 1 //Deklarasi

	// for counter <= 5{ //Kondisi
	// 	fmt.Println(counter)
	// 	counter++ //Increment
	// }

	// for i := 1; i <= 5; i++{
	// 	for j := 1; j<i; j++{
	// 		fmt.Print(j)
	// 	}
	// 	fmt.Println(i)
	// }

	slice := []string{"Fatur", "Udin", "Jajang"}

	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	for i, value := range slice{
		fmt.Println("Index", i, "=", value)
	}
}