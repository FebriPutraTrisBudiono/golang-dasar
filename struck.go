package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// STRUCK
	// var febri Customer
	// febri.Name = "Febri Putra"
	// febri.Address = "Jember"
	// febri.Age = 25
	// fmt.Println(febri)

	// STRUCT LITERALS
	// febri := Customer{
	// 	Name:    "Febri Putra",
	// 	Address: "Jember",
	// 	Age:     25,
	// }
	// fmt.Println(febri)

	// STRUCT LITERALS V2
	febri := Customer{"Febri Putra", "Jember", 25}
	fmt.Println(febri)
}
