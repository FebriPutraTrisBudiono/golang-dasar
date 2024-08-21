package main

import "fmt"

func main() {
	// names := [...]string{"Febri", "Putra", "Tris", "Budiono"}
	// slice := names[2:4]

	// fmt.Println(slice[0])
	// fmt.Println(slice[1])

	// APPEND SLICE
	// days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	// daySlice1 := days[5:]
	// daySlice1[0] = "Sabtu Baru"
	// daySlice1[1] = "Minggu Baru"
	// fmt.Println(days)

	// daySlice2 := append(daySlice1, "Libur Baru")
	// daySlice2[0] = "Ups"
	// fmt.Println(daySlice2)
	// fmt.Println(days)

	// MAKE SLICE
	// newSlice := make([]string, 2, 5)
	// newSlice[0] = "Febri"
	// newSlice[1] = "Febri123"

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// COPY SLICE
	// fromSlice := days[:]
	// toSlice := make([]string, len(fromSlice), cap(fromSlice))

	// copy(toSlice, fromSlice)

	// fmt.Println(toSlice)

	// ARRAY VS SLICE
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
