package main

import "fmt"


// PENGGUNAAN DEFER FUNCTION 

func logging(){
	fmt.Println("Function telah terpanggil")
}

func app(){
	defer logging()
	fmt.Println("Memanggil fungsi logging")
}

func appDone(){
	message := recover()
	fmt.Println("TERJADI ERROR", message)
	fmt.Println("APLIKASI SELESAI")
}

func appWorking(error bool){
	defer appDone()
	if error{
		panic("APLIKASI TIDAK BERJALAN")
	}
	fmt.Println("APLIKASI BERJALAN")
}

func main() {
	app()
	appWorking(true)
}