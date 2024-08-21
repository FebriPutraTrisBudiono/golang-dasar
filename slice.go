package main

import "fmt"

func main() {
	// names := [...]string{"Febri", "Putra", "Tris", "Budiono"}
	// slice := names[2:4]

	// fmt.Println(slice[0])
	// fmt.Println(slice[1])

	// Append slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)
}
