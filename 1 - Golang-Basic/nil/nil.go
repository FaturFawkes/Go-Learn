package main

import "fmt"

func NewMap(name string) map[string]string{
	if name == ""{
		return nil
	} else {
		return map[string]string{
			"name" : name,
		}
	}
}

func main() {
	// var person map[string]string = nil
	var person map[string]string = NewMap("fatur")
	// var fatur = NewMap("fatur")
	fmt.Println(person)
}