package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Febri",
		"address": "Jember",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Febri Putra"
	book["wrong"] = "Ups"

	delete(book, "wrong")

	fmt.Println(book)
}
