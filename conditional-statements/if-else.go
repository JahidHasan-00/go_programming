package main

import "fmt"

func main() {
	/**If-Else Statement using operators:
	   // >, <, >=, <=, ==, !=
	   // and => &&, or => ||
	   // not = !
	**/
	num := 18

	if num > 18 {
		fmt.Println("You are eligible to marry")
	} else if num == 18 {
		fmt.Println("You should wait for crossing 18")
	} else {
		fmt.Println("You are not eligible to marry")
	}
	//Output: You should wait for crossing 18
	//Another Example:
	if num >= 18 {
		fmt.Println("You are eligible to vote")
	} else if num <= 17 {
		fmt.Println("You are not eligible to vote")
	} else {
		fmt.Println("You should wait for crossing 18")
	}

	//Another Example of ( &&, || ) operator:
	a := 10
	b := "Jahid"
	if a == 10 && b == "Jahid" {
		fmt.Println("Both are correct")
	} else {
		fmt.Println("Both are not correct")
	}
	//
	if a > 10 && b == "Jahid" {
		fmt.Println("Both are correct")
	} else {
		fmt.Println("Both are not correct")
	}
	//
	if a > 10 || b == "Jahid" {
		fmt.Println("Both are correct")
	} else {
		fmt.Println("Both are not correct")
	}

	//Another Example of ( ! ) operator:
	x := true
	fmt.Println(!x) // false
	x = false
	fmt.Println(!x) // true

	//Switch Statement:
	//Switch statement is used to execute a block of code among many alternatives.

	c := 3
	switch c {
	case 1:
		fmt.Println("a is 1")
	case 2, 3:
		fmt.Println("a is either 2 or 3")
	default:
		fmt.Println("a is neither 1 nor 2, 3")
	}
}
