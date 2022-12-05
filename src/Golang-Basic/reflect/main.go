package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `json:"name" max:"10" required:"true"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value :=  reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			} 
		}
	}
	return true
}

func main() {
	sample := Sample{"Fatur"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0).Name
	fmt.Println(structField)

	// Get data from struct tag
	required := sampleType.Field(0).Tag.Get("max")
	fmt.Println(required)

	// Struct tag cheking
	value := isValid(sample)
	fmt.Println(value)
}