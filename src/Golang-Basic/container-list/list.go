package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Nur")
	data.PushBack("Fatur")
	data.PushFront(1)
	f := data.Front()
	fmt.Println(f.Value)

	// Iterasi dar depan ke belakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}