package main

import (
	"Golang-Basic/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}