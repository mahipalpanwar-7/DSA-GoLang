// Space Complexity

/*

   space complexity = Auxiluary space + Input space

   Auxiliary space = the space that we takes to solve the problem
   Input Space  =  the space that we takes to store the input

   calculated in Big-oh notation

*/

package main

import "fmt"

func main() {
	a := 2 // here a and b is input space  and C is auxiliary
	b := 3
	c := a + b
	fmt.Println(c)    // Space Complexity is O(3) for this example but we can ignore constants then it will be O(1) for N number of array input O(N)
}
