package main

import "fmt"

func main() {
	// index array
	// var names [3]string
	// names[0] = "Febri"
	// names[1] = "Putra"
	// names[2] = "Tris"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	// array langsung
	// var values = [3]int{
	// 	90,
	// 	80,
	// 	95,
	// }
	// fmt.Println(values)

	// function array
	var values = [...]int{
		90,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
