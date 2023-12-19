package main

import "fmt"

func main()  {
	// variable harus di gunakan atau jadi error
	// var name = "Ambar Dwi Saputra"
	// := hanya deklarasi awal
	// variable harus di deklarasikan tipe datanya, kecuali langsung di assign (go dapat mendeteksi)
	name := "Ambar Dwi Saputra"
	fmt.Println(name)

	name = "Ambar D S"
	fmt.Println(name)

	// multiple variable
	var (
		firstName = "Ambar"
		midleName = "Dwi"
		lastName = "Saputra"
	)

	fmt.Println(firstName)
	fmt.Println(midleName)
	fmt.Println(lastName)
}