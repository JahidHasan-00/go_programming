package main

import "fmt"

var (
	arr1 = [3]string{"I", "Love", "You"}
)

func main() {
	arr := [5]int{4, 2, 1, 3, 5}

	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr1[1])
}

/**
 #Compilation:
	Code Segment:
	main = func() {....}

**/
