package main

import "fmt"

func main() {
	//pointer or memory address
	x := 20
	fmt.Println(x)

	p := &x //ampersand & => address of
	*p = 40

	fmt.Println(x)
	// fmt.Println("Address:", p)               //now p is the address of x
	// fmt.Println("Value at the address:", *p) // ( * ) value of address
}
