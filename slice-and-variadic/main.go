package main

import "fmt"

func main() {
	// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(arr)

	// // slice from existing array
	// s := arr[1:5]
	// fmt.Println(s)      //[2 3 4 5]
	// fmt.Println(len(s)) // 4
	// fmt.Println(cap(s)) // 9

	// // slice from a slice
	// s1 := s[1:3]
	// fmt.Println(s1)      //[3 4]
	// fmt.Println(len(s1)) // 2
	// fmt.Println(cap(s1)) // 8

	// slice literal:
	// sl := []int{1, 2, 5}                                        //pointer, length, capacity
	// fmt.Println("slice:", sl, "len:", len(sl), "cap:", cap(sl)) // [1 2 5] len: 3 cap: 3

	sm := make([]int, 3) // [0 2 0] len:3 cap:3
	sm[1] = 2

	fmt.Println(sm)      // [0 2 0]
	fmt.Println(len(sm)) // 3
	fmt.Println(cap(sm)) // 3.
}

/**
	Compilation Time:
	 #Code Segment:

	 main = func (){....}

**/
/***
	1. slice from existing array
	2. slice from slice
	3. slice literal
	4. slice make()
***/
