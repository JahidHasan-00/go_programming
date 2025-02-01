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

func main() {
	result := add(10, 20)

	//Anonymous Function
	//IIFE (Immediately Invoked Function Expression)
	func(res int) {
		fmt.Println("Sum of those values: ", res)
	}(result)
}

func init() {
	fmt.Println("I'll be execute first")
}
