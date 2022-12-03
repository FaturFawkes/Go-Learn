package main

import "fmt"

func endApp() {
	// RECOVER dijalankan di luar scope panic (di defer funtion)
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error ", message)
	}
	fmt.Println("Aplikasi selesai")
}


func runApp(error bool) {
	defer endApp()
	fmt.Println("Aplikasi mulai")
	if error {
		// PANIC
		panic("ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}


func main() {
	runApp(true)
}