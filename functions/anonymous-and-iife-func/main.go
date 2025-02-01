/**Anonymous and IIFE Function in Golang:
   * Anonymous function is a function defined without a name.
   * IIFE (Immediately Invoked Function Expression) is a function that runs as soon as it is defined.
**/

package main

import "fmt"

//Named or Standard Function
func add(x, y int) int {
	return x + y - x*y
}

//Function returns an anonymous function
func multiplier(factor int) func(int) int {

	return func(a int) int {
		return a * factor
	}
}

func main() {
	result := add(10, 20)

	result1 := multiplier(5)
	fmt.Println(result1(5))

	//Anonymous Function
	//IIFE (Immediately Invoked Function Expression)
	func(res int) {
		fmt.Println("Sum of those values: ", res)
	}(result)

	//Assigning an anonymous function to a variable
	greet := func(name string) {
		fmt.Println("Hello, ", name)
	}
	greet("John")
	greet("Doe")
}

func init() {
	fmt.Println("I'll be execute first")
}
