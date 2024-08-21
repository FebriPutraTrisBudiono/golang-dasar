package main

import "fmt"

func main() {
	// BREAK
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}

	// 	fmt.Println("Perulangan ke ", i)
	// }

	// CONTINUE
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke ", i)
	}
}
