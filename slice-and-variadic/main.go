package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	//slice from existing array
	s := arr[1:5]
	fmt.Println(s)      //[2 3 4 5]
	fmt.Println(len(s)) // 4
	fmt.Println(cap(s)) // 9

	//slice from a slice
	s1 := s[1:3]
	fmt.Println(s1)      //[3 4]
	fmt.Println(len(s1)) // 2
	fmt.Println(cap(s1)) // 8
}

/**
	Compilation Time:
	 #Code Segment:

	 main = func (){....}

**/
