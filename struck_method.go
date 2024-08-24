package main

import "fmt"

type User struct {
	Name string
}

func (user User) sayHello() {
	fmt.Println("Hello my name is", user.Name)
}

func main() {
	febri := User{Name: "Febri"}
	febri.sayHello()
}
