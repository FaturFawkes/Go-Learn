package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()

	// fmt.Println(now)

	// Mmebuat time sendiri
	// utc := time.Date(2022, 12, 04, 19, 39, 00, 000, time.Local)
	// fmt.Println(utc)

	// Parsing time
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-12-04")
	fmt.Println(parse)
}