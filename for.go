package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// FOR DENGAN STATEMENT
	// for counter := counter; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke ", counter)
	// }

	// FOR RANGE
	names := []string{"Febri", "Putra", "Tris"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
