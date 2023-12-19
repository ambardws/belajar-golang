package main

import "fmt"

func main()  {
	// constant walaupun tidak digunakan saat di run tidak error
	// const tidak bisa di assign lagi

	const (
		firstName string = "Ambar"
		lastName = "Dwi Saputra"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	
}