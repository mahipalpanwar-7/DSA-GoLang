/* time complexity Notations

Big-Oh (O) Represents := Worst Case, Meaning := Upper Bound
Theta (Θ)  Represents := Average/Exact Growth Case, Meaning := Upper Bound
omega (Ω)  Represents := Worst Case , Meaning := Upper Bound

*/


//rate of which time taken increases with respect to the input size


/* STEPS 

1. Time Complexity to be computed in WORST CASE scenario
2. Avoid Constants 
3. Avoid Lower Values

*/


package main 

import "fmt"


func main(){
 
// first example 

for i:=1;i<=5;i++{
	fmt.Println("Golang")   // time complexity O(5*3)  
}
	
// if we take N then this will be O(N*3)     final result according to the convention (avoid lower values) O(N)

/* number of steps this code will take will come inside O(here) Big-Oh
increment -> check -> print (mainly three steps) */


// second example 


marks := 85
if (marks <25){
	fmt.Println("grade D")
}else if (marks<45){
	fmt.Println("grade C")
}else if(marks<65){
	fmt.Println("grade B")
}else{
	fmt.Println("grade A")
}

 /* time complexity  
 
 Best case O(2)    = check->print
 Worst case O(4)   = check->check->check->print

*/

// Third example 

N:=10
for i:=0;i<=N;i++{
	for j:=1;j<=N;j++{    
       fmt.Println("")
	}
}

/* 
i=0 {j=0,1,2,3...N}
i=1 {j=0,1,2,3...N}
i=N {j=0,1,2,3...N}

every time it is running for (N+N+N+N+N+.....N) times 

how many times its iterating  O(N*N) = O(N²)
*?

// if for j:=1;j<=i;j++ then

/*
i=0 {j=0}
i=1 {j=0,1}
i=N {j=0,1,2,3...N}

number of iterations = (1+2+3+4....+n)  => N*(N+1)/2 => N²/2+N/2 => according to the rules end result will be => N²/2  => O(N²)

*/


}