package main

import "fmt"

func random() interface{} {
	return false
}

func catch(){
	if r := recover(); r != nil{
		fmt.Println("app error")
	} 
}

func main() {
	defer catch()
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int: 
		fmt.Println("Integer", value)
	default:
		fmt.Println("unknown")
	} 
	// var result interface{} = random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	
}
