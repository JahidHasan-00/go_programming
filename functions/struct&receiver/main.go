package main

import "fmt"

type User struct {
	//properties / member variable:
	Name string
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

//Receiver Function:
func (usr User) printDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

//

func main() {
	var user1 User
	user1 = User{ //Instantiate
		Name: "Jahid",
		Age:  26,
	}
	// printUserDetails(user1)
	user1.printDetails()
	user1.call(5)

	user2 := User{ // Instance or Object
		Name: "Habib",
		Age:  30,
	}
	printUserDetails(user2)
}

/****
	Two phases:
	 1.Compilation phase (Create a binary executable file during compilation phase)
	 2.Execution phase

	##Compilation:
	   User = type User struct{.....}
	   printUserDetails = func(){....}
	   printDetails = func(){....} //User
	   call = func(){....}
	   func main(){....}

***/
