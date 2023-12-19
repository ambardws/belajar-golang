package main

import "fmt"

func main() {
	type noKTP string

	var ktpAmbar noKTP = "123455666"
	var contoh string = "222222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpAmbar)
	fmt.Println(contohKTP)

}