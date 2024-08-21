package main

import "fmt"

func main() {
	name := "Putra"

	// if name == "Febri" {
	// 	fmt.Println("Hello febri")
	// } else if name == "Putra" {
	// 	fmt.Println("Hello Putra")
	// } else {
	// 	fmt.Println("Hi, boleh kenalan ?")
	// }

	// SHORT STATEMENT
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
