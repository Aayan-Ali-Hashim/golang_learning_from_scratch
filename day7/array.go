package main

import "fmt"


func main() {
	var a [5]int // This declares an array of 5 integer

	// In GO we need to mention the size of array so not like Python 


	fmt.Println("Initial array:", a)
	
	// Assigning values to the array
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50

	fmt.Println(a)
	
	// creating array in one line 
	
	b := [5]int{10, 20, 30, 40, 50}
	
	fmt.Println(b)
}




