package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := 0; i < value; i++ {
		result *= i
	}
	return value
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialRecursive(10))
}
