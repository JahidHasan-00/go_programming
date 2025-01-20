package main

import "fmt"

// x := 1;

const myVar1 int = 100
const PI = 3.1416

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

	//Go Variable Declaration in a Block
	var (
		myDailyUsedPro     = 5
		myDailyWalking     int
		myDailyUsedProName = "Laptop, Mobile, T-Shirt"
	)

	fmt.Print("\n")
	fmt.Println(myDailyUsedPro)
	fmt.Println(myDailyWalking)
	fmt.Println(myDailyUsedProName)

	/**Constants: Unchangeable and Read-only
	   ---> When a constant is declared, it is not possible to change the value later
	**/
	const A1 = 10.5
	// A1 = 10 //Can not Assign 10. When a constant is declared, it is not possible to change the value later
	// fmt.Println(A1) //Output: Error
	fmt.Println(myVar1) //myVar1 declared outside of top of the function
	fmt.Println(PI)     //PI declared outside of the function

	// Multiple Constants Declaration:
	const (
		phoneName = "iPhone"
		brand     = "Apple"
		cost      = 120000
	)

	fmt.Println(phoneName, brand, cost)

	//If we want to add a space between string arguments, we need to use " " inside fmt.Print() function:
	const i, j = 880, "Bangladesh"
	fmt.Print(i, " ", j)
	fmt.Printf("\n")
	/**The Printf() Function:
	   --> The Printf() function first formats its argument based on the given formatting verb and then prints them.
	   -->Here we will use two formatting verbs:
			%v is used to print the value of the arguments
			%T is used to print the type of the arguments
	**/
	fmt.Printf("i has value of : %v and i has type of : %T \n", i, i)
	fmt.Printf("j has value of : %v and j has type of : %T \n", j, j)
}
