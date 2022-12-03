package main

import "fmt"

func logging() {
	fmt.Println("Function selesai dijalankan")
}

// DEFER
func runApp(value int) {
	defer logging() // akan menjalankan func logging sebelum function runApp berakhir dan tidak peduli jika ada error
	fmt.Println("Aplikasi berjalan")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApp(10)
}