package main

import "fmt"

//Pass By Value
//Pass By Reference

type User struct {
	Name    string
	Age     int
	Salary  float32
	favFood [3]string
}

func printInstance(profile *User) {

	fmt.Println(profile)
	fmt.Println(profile.Name)

}

func print(numbers *[3]int) {

	fmt.Println(numbers[1])

}

func main() {
	arr := [3]int{4, 5, 6}
	print(&arr)

	person := User{
		Name:   "Habibur",
		Age:    30,
		Salary: 500000,
	}

	p := &person

	printInstance(p)
	printInstance(&person)
}

/***
	##Compilation Phase:

		Code Segment:
			User = type User Struct
			printInstance = func(){....}
			print = func (){...}
			main = func(){...}

***/
