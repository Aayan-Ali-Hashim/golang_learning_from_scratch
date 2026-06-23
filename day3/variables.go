package main
import "fmt"

func main() {

	// Three different techniques to create variables
	// You can use var to create variable
	// You can define a type of variable but it is optional, like Python; Go can auto assign datatype to variables
	var a = "hello"
	var b int = 2
	var c float64 = 1.0
	// You can use the shorthand := operator to create variable without using any var keyword
	d := "go is great"
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
	
}


