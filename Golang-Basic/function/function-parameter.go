package main

import "fmt"

func main() {
	sayHello("Nur", "Fatchurohman")
	fmt.Println(getHello("Fatur"))
	fmt.Println(getPujian("fatur"))

	angka1, angka2 := angka()
	fmt.Println(angka1, angka2)
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// FUNCTION RETURN VALUE

func getHello(nama string) string {
	return "Hello " + nama
}

// FUNCTION RETURN VALUE MULTIPLE

func getPujian(nama string) (string, string) {
	return "Anjay kamu ganteng sekali " + nama + " aku jadi iri", "sama kamu"
}

func angka() (int, int) {
	angka1 := 1
	angka2 := 2
	result := angka1 + angka2

	return result, angka2
}
