package main

import "fmt"

func getGoodBye(name string) string {
	return "Selamat Tinggal " + name
}

func main() {
	result := getGoodBye
	fmt.Println(result("Febri"))
}
