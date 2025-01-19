package main

import "fmt"

// x := 1;

func main() {
	/**
	int
	float32
	string
	bool
	**/
	//Variable Declaration With Initial Value
	var a int = 20      //type is int
	var b string = "50" //type is String
	var c = true        //type is inferred
	d := false          //type is inferred

	fmt.Println(a)
	fmt.Print(b)
	fmt.Println(50)
	fmt.Println(c)
	fmt.Println(d)

	/**Variable Declaration Without Initial Value:
	   --> In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type
	**/
	var a1 string
	var b1 int
	var c1 bool

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)

	/**Value Assignment After Declaration:
	  --> It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.
	*/

	var student string
	student = "Jhon"

	fmt.Println(student, "\n")

	// //Since := is used outside of a function, running the program results in an error:
	// fmt.Println(x) //.\variable-datatype.go:6:1: syntax error: non-declaration statement outside function body

	/**Go Multiple Variable Declaration:
	  --> In Go, it is possible to declare multiple variables on the same line.
	  -- example below:
	*/
	var num1, num2, num3, num4 int = 9, 8, 7, 6 //int --> type
	fmt.Println(num1, num2, num3, num4)

	//Note: If you use the type(int, float32, string, bool) keyword, it is only possible to declare one type of variable per line.
	// var z string = true;

	//If the type keyword is not specified, you can declare different types of variables on the same line
	var myAge, myName, programmer = 26, "Jahid Hasan", true
	fmt.Println(myAge, myName, programmer)

}

// func main() {
// 	fmt.Print("Hello World")
// }
