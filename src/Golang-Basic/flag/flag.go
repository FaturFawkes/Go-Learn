package main

import (
	"flag"
	"fmt"
)

func main() {
	// Untuk membuat flag sendiri ketika akan run file golang
	host := flag.String("host", "root", "Put your host")
	user := flag.String("user", "root", "Put your user")
	password := flag.String("password", "root", "Put your password")

	flag.Parse()

	fmt.Println(*host, *user, *password)
}