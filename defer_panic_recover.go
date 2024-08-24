package main

import (
	"fmt"
)

// DEFER
// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }

// func runApplication() {
// 	defer logging()
// 	fmt.Println("Run Application")
// }

// PANIC
func endApp() {
	fmt.Println("End App")
	// RECOVER
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	// DEFER
	// runApplication()

	// PANIC
	runApp(true)
}
