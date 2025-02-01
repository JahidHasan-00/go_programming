package main

import "fmt"

func printMsgWelcome() {
	fmt.Println("Welcome to the Application")
}

func summation() int {
	a := 10
	b := 20
	sum := a + b // Summation of a and b
	return sum
}

func subtraction() int {
	a := 10
	b := 20
	sub := b - a // Subtraction of a and b
	return sub
}

/**Function with Return Values and Types:
   --> Function with return values and types are used to return a value from a function.
**/
//Single return value function
func multiplication() int { // a = 10, b = 20
	var a int = 40
	var b int = 20
	return a * b // Multiplication of a and b

}

//Multiple return value function
func getNumbers() (int, int) {
	a := 10
	b := 20
	mod := b % a
	div := b / a
	return mod, div
}

//String return value function
func getMyFullName(firstName, lastName string) string {
	fullName := firstName + " " + lastName
	fmt.Println("Full Name: ", fullName)
	return fullName
}

func userInput() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func calculate() int {
	var num1, num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	return num1 + num2
}

//Display the common output of the function
func displayOutput(add int, sub int, mul int, mod int, div int, fullName string) {
	fmt.Println("Addition: ", add)
	fmt.Println("Subtraction: ", sub)
	fmt.Println("multiplication: ", mul)
	fmt.Println("Modulas: ", mod)
	fmt.Println("Division: ", div)
	fmt.Println("Full Name: ", fullName)
}

//Display get user name output function
func display(name string, sum int) {
	fmt.Println("Hello, ", name)
	fmt.Println("Summation: ", sum)
}

func goodBye() {
	fmt.Println("Thank you for using the application")
	fmt.Println("Good Bye")
}

//Main Function:
func main() {
	printMsgWelcome()
	add := summation()
	//Now calculate the subtraction operation from another function
	sub := subtraction()
	//Calculate the multiplication operation
	mul := multiplication()
	mod, div := getNumbers()
	fullName := getMyFullName("Jahid", "Hasan")
	displayOutput(add, sub, mul, mod, div, fullName)
	//Get user name as input
	name := userInput()
	sum := calculate()
	display(name, sum)
	goodBye()
}
