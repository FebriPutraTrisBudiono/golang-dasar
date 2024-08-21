package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Febri"
	middleName = "Putra"
	lastName = "Tris"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
