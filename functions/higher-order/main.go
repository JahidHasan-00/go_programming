package main

import "fmt"

//Returns a function
func call() func(num1 int, num2 int) {
	return mul
}
func mul(num1, num2 int) {
	var mul = num1 * num2
	fmt.Println(mul)
}

func main() {
	res := call()
	res(2, 5)
}
