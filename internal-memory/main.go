package main

import "fmt"

const a = 10 // constant value
var p = 100  // variable value

func call() {
	add := func(x, y int) {
		z := x + y
		fmt.Println(z)
	}
	add(5, 6)
	add(p, a)
}

func main() {
	call()
	fmt.Println(a)
}

func init() {
	fmt.Println("Hello")
}

/**
Golang have 2 phases:
	1. Compilation phase
	2. Execution phase


	*********** Compilation phase **********

	1. It will check for syntax errors
	2. It will check for type errors
	3. It will check for variable declaration
	4. It will check for function declaration
	5. It will check for package declaration
	6. It will check for import declaration
	7. It will check for constant declaration

	**Code Segment: **
		a = 10 // constant value is constant thats why it will be stored in code segment
		call(){....}
		add(){....}
		main(){....}
		init(){....}

	************ Execution phase ***********



**/
