package main

import "fmt"

type Filter func(string)string

func filterWords(words string, filter Filter){
	fmt.Println("Kamu bilang", filter(words))
}

func condition(words string)string {
	if words == "anjing"{
		return "kasar"
	} else {
		return words
	}
}

func main(){
	filterWords("anjing", condition)
	filterWords("fatur ganteng", condition)
}