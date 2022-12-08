package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fatur")
	if result != "Hi Fatur" {
		// t.Fail now tidak akan menjalankan program dibawahnya	
		// t.Error bisa menambahkan informasi gagalnya dan otomatis menjalankan t.fail()
		t.Error("Error Hi Fatur")
		t.Fail()
	}
	fmt.Println("Testing Fatur Done")
}

func TestHelloWorldRohman(t *testing.T) {
	result := HelloWorld("Rohman")
	if result != "Hello Rohman" {
		panic("Result is not Hello Rohman")
	}
	fmt.Println("Testing Rohman Done")
}