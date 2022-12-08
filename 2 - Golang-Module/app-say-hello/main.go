package main

import (
	"fmt"

	"github.com/FaturFawkes/go-say-hello/v2"
)

func main() {
	res := go_say_hello.SayHello("fatur", "Bogor")
	fmt.Println(res)
}