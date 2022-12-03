package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error){
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		result := nilai/pembagi
		return result, nil
	}
}

func main() {
	hasil, err := pembagian(2, 0)
	if err == nil{
		fmt.Println("Hasil = ", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}