package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Jember", "Jawa Timur", "Indonesia"}
	address2 := &address1

	address2.City = "Surabaya"

	// Tanpa Operator *
	// address2 = &Address{"Solo", "Jawa Tengah", "Indonesia"}
	// Dengan Operator *
	*address2 = Address{"Solo", "Jawa Tengah", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
