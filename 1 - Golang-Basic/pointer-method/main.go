package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {
	fatur := Man{"Fatur"}
	fatur.Married()
	fmt.Println(fatur.Name)
}