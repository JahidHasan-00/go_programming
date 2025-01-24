package main

import (
	"fmt"
	"mathlib/mathlib"
)

// Global Scope
var (
	num1 = 10
	num2 = 20
)

func main() {
	fmt.Println("This is the main function body and this is the local scope")
	add(num1, num2)
	fmt.Println("Showing Custom Package")
	mathlib.Multiplication(num1, num2)
}
