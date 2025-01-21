package main

import "fmt"

func subtraction(a int, b int) {
	sub := b - a // Subtraction of a and b
	fmt.Println("Subtraction: ", sub)
}

func main() {
	a := 10
	b := 20
	sum := a + b // Summation of a and b
	fmt.Println("Summation: ", sum)
	//Now calculate the subtraction operation from another function
	subtraction(a, b)
}
