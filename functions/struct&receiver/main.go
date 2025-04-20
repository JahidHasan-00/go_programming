package main

import "fmt"

type User struct {
	//properties:
	Name string
	Age  int
}

func main() {
	var user1 User
	user1 = User{ //Instantiate
		Name: "Jahid",
		Age:  26,
	}

	user2 := User{ // Instance or Object
		Name: "Habib",
		Age:  30,
	}

	fmt.Println("Name:", user1.Name, "\n", "Age:", user1.Age)
	fmt.Print("\n")
	fmt.Println("Name:", user2.Name, "\n", "Age:", user2.Age)
}

/****
	Two phases:
	 1.Compilation phase (Create a binary executable file during compilation phase)
	 2.Execution phase

	##Compilation:
	   User = type User struct{.....}
	   func main(){....}

***/
