package main

import "fmt"

func Ups() interface{} {
	// return "Ups"
	// return 1
	return true
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
