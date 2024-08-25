package main

import "fmt"

func random() interface{} {
	return 1
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) //panic
	// fmt.Println(resultInt)

	// Type Assertions Switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Uknown")
	}
}
