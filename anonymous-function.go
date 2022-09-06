package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked!", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		// fmt.Println(name == "user")
		return name == "user"
	}

	registerUser("fatur", blacklist)
	registerUser("admin", blacklist)
}