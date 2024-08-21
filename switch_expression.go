package main

import "fmt"

func main() {
	name := "Febri"

	// switch name {
	// case "Febri":
	// 	fmt.Println("Hello Febri")
	// case "Putra":
	// 	fmt.Println("Hello Putra")
	// default:
	// 	fmt.Println("Hi, boleh kenalan ?")
	// }

	// SHORT STATEMENT
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }

	// TANPA KONDISI
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
