package main

import (
	"fmt"
	"mathlib/mathlib"
)

// Global Scope
var (
	num1 = 10
	num2 = 20
	num3 = 30
)

// const num3 = 20

//   num3 = 30 // This will throw an error because we cannot reassign a value to a constant

func main() {
	num2 = 10
	var num3 = 40 // Local Scope variable num3
	fmt.Println("This is the main function body and this is the local scope")
	add(num1, num2)
	fmt.Println("Showing Custom Package")
	mathlib.Multiplication(num1, num2)
	getValue(num3)

	// Block Scope
	if num1 == num2 {
		fmt.Println("num1 is equal to num2")
	} else if num2 == num3 {
		fmt.Println("num2 is equal to num3")

	} else {
		x := 10
		fmt.Println("num1 is not equal to num2")
		fmt.Println("But x is equal to num1: ", x, "=", num1)
	}
	// fmt.Println(x); // Block scope variable x is not accessible here
}

func getValue(x int) {
	fmt.Println(x)
}
