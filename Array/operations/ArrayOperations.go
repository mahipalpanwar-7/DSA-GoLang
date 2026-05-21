package main

import "fmt"

func main() {

	// array traversal
	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

	}

	// length of array
	fmt.Println("length of array", len(arr))

	// copy array
	var ar [4]int
	ar = arr
	fmt.Println(ar)

	// Comparison  of array

	fmt.Println("check Array ar and arr are same", ar == arr)
	fmt.Println("check Array ar and arr are not same", ar != arr)

	// 2D array

	arrr := [4][3]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}
	fmt.Println(arrr)

	//traversal of 2D array

	for i := 0; i < len(arrr); i++ {
		for j := 0; j < len(arrr[i]); j++ {
			fmt.Print(arrr[i][j], " ")
		}
		fmt.Println()
	}
}
