package main

import "fmt"

func getFullName() (string, string) {
	return "Febri", "Putra"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	// MENGHIRAUKAN RETURN VALUE
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
