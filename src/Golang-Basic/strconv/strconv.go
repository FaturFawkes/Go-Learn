package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, _ := strconv.ParseBool("true")
	fmt.Println(boolean)
	convNum, err := strconv.ParseInt("10000", 10, 64)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(convNum)

}