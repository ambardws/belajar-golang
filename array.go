package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Ambar"
	names[1] = "Dwi"
	names[2] = "Saputra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// var values = [...] {} sesuai panjangnya data yang di declare
	var values = [3]int {
		1,
		2,
		3,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}