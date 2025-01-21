package main

import "fmt"

func subtraction(a int, b int) {
	sub := b - a // Subtraction of a and b
	fmt.Println("Subtraction: ", sub)
}

/**Function with Return Values and Types:
   --> Function with return values and types are used to return a value from a function.
**/
//Single return value function
func multiplication(a int, b int) int { // a = 10, b = 20
	mul := a * b
	return mul
}

//Multiple return value function
func getNumbers(a int, b int) (int, int) {
	mod := a % b
	div := a / b
	return mod, div
}

//String return value function
func getMyName(name string) {
	fmt.Println("My Name is : ", name)
}

//
func getMyFullName(firstName, lastName string) string {
	fullName := firstName + " " + lastName
	return fullName
}

func main() {
	a := 10
	b := 20
	sum := a + b // Summation of a and b
	fmt.Println("Summation: ", sum)
	//Now calculate the subtraction operation from another function
	subtraction(a, b)
	subtraction(30, 50)
	//Calculate the multiplication operation
	mul := multiplication(a, b)
	fmt.Println("multiplication: ", mul)
	//Another Example:
	mod, div := getNumbers(a, b)
	fmt.Println("Modulas: ", mod)
	fmt.Println("Division: ", div)
	getMyName("Jahid")
	fullName := getMyFullName("Jahid", "Hasan")
	fmt.Println("Full Name: ", fullName)
}
