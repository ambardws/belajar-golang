package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c) // 20
	c += 5
	fmt.Println(c) // 25
	c++
	fmt.Println(c) // 26
}