package main

import "fmt"

func changeSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

func print(a ...int) {
	fmt.Println(a)
}

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

	// sm := make([]int, 3) // [0 2 0] len:3 cap:3
	// sm[1] = 2

	// fmt.Println(sm)      // [0 2 0]
	// fmt.Println(len(sm)) // 3
	// fmt.Println(cap(sm)) // 3.

	//Another way to use make function:
	// sm1 := make([]int, 4, 5) // [0 0 0 0] len: 4, cap: 5

	// sm1[1] = 5
	// sm1[0] = 2
	// sm1[3] = 6

	// fmt.Println(sm1)      // [2 5 0 6], pointer: indexOf-0 of memory address
	// fmt.Println(len(sm1)) // len: 4
	// fmt.Println(cap(sm1)) // cap: 5

	// sm2 := sm1[1:4] //[5 0 6]

	// fmt.Println(sm2)      // [5 0 6]
	// fmt.Println(len(sm2)) // 3
	// fmt.Println(cap(sm2)) // 4

	// Empty slice or nil slice:
	// var x []int      // pointer = nil, length = 0, cap = 0
	// x = append(x, 1) // [1], len = 1, cap = 1
	// x = append(x, 2) // [1, 2], len = 2, cap = 2
	// x = append(x, 3) // [1, 2, 3], len = 3, cap = 4

	// y := x // [1, 2, 3], len = 3, cap = 4

	// x = append(x, 4)
	// y = append(y, 5)

	// x[0] = 10

	// fmt.Println(x, len(x), cap(x)) //[1 2 3 5], len = 4, cap = 4
	// fmt.Println(y, len(y), cap(y)) //[1 2 3 5], len = 4, cap = 4
	// // fmt.Println(len(x))    //
	// // fmt.Println(cap(x))    // 3

	//Another Example:
	x1 := []int{1, 2, 3, 4, 5}
	x1 = append(x1, 6)
	x1 = append(x1, 7)

	a := x1[4:]
	y1 := changeSlice(a)

	fmt.Println(x1)      // [1 2 3 4 10 6 7]
	fmt.Println(y1)      // [10 6 7 11]
	fmt.Println(x1[0:8]) // [1 2 3 4 10 6 7 11]

	//Pass argument for variadic function:
	print(5, 10, 12)

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
	4. slice make() with length
	5. slice make() with length and capacity
	6. Empty slice or nil slice using append() function
***/
