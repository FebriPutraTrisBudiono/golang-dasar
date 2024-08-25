package main

import "fmt"

type Address struct {
	Address, City, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// Pass by value
	// address2 := address1
	// Pass by reference dengan pointer
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

}
