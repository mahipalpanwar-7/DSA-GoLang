// basic array program

package main

import "fmt"

func main() {
	// array
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// invalid operation for array a:=append(arr,1,2,3); fmt.Println("Array",a)

	var ar [3]int
	ar[0] = 0
	ar[1] = 1
	ar[2] = 2
	fmt.Println("array1", arr)
	fmt.Println("array2", ar)

	//slice
	slice := []int{0, 1, 2, 3}
	s := append(slice, 4, 5, 6)

	fmt.Println("slice1", slice)
	fmt.Println("slice2", s)

	slice2 := make([]int, 3, 5)
	/* slice capacity Overflow
	Creates new bigger array
	Copies old values
	Appends new values
	Returns new slice */
	slice2 = append(slice2, 1, 2, 3)
	fmt.Println("slice1", slice2)

	// slicing of underlying array
	slice3 := arr[1:5]
	fmt.Println("slice3", slice3)
}
