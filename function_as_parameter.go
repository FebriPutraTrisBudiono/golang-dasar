package main

import "fmt"

type Filter func(string) string

func sayWithHello(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func filter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayWithHello("Anjing", filter)

	result := filter
	sayWithHello("Febri", result)
}
