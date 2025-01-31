/**init function: An init function is a function used to initialize the application state before the main function execution. The init function in Go is a special function that is automatically executed to initialize a package

	* It takes no arguments and returns no result
	* Runs after global variable initialization in the package but before the main function (if in the main package)
	* Guaranteed to execute exactly once per package, even if the package is imported multiple times
	* A package can have multiple init functions
	* Within a single file, they execute in the order of declaration
	* Across files in a package, execution order depends on the lexical file order (determined by the compiler, but not recommended to rely on)
	* Dependencies are initialized first. For example:
		 If package A imports package B, then B's init function is executed before A's
	* Execution sequence:
		1. Imported packages' `init` functions
		2. Current package's global variables
		3. Current package's `init` function
		4. main function (if applicable)
**/

package main

import "fmt"

var globalVar = initVar()

func init() {
	fmt.Println("First Init (uses globalVar =", globalVar, ") ")
}

func init() {
	fmt.Println("Second init")
}

func main() {
	fmt.Println("Main function")
}

func init() {
	fmt.Println("Third init")
}

func initVar() int {
	fmt.Println("Global Variable Initialize")
	return 50
}
