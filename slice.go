package main

import (
	"fmt"
)

func main() {
	names := [...]string {
		"Ambar", "Dwi", "Saputra", "Joko", "Budi", "Nugraha",
	}
	// slice tidak perlu menentukan jumlah array
	// var slice []string =
	slices := names[4:6]
	slices2 := names[:4]
	slices3 := names[3:]

	fmt.Println(slices[0])
	fmt.Println(slices[1])
	fmt.Println(slices2)
	fmt.Println(slices3)
	fmt.Println(len(slices))
	fmt.Println(cap(slices))

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ambar"
	newSlice[1] = "Dwi Saputra"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}