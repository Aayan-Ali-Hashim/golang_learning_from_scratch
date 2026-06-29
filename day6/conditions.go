package main

import "fmt"

// Conditions in GO are easy and just like Python

func main() {
	name := "Alice"
	if name == "Alice" {
		fmt.Println("Hello, Alice!")
	} else if name == "Bob" {
		fmt.Println("Hello, Bob!")
	} else {
		fmt.Println("Hello, stranger!")
	}

	// You can also declare a variable in the if statement itself

	if number := 10; number > 5 {
		fmt.Println("The number is greater than 5.")
	}

	// && and || operators are also available in GO for comparison

	// Below is an example

	if name == "Alice" && len(name) > 0 {
		fmt.Println("Hello, Alice! Your name is not empty.")
	}
	if name == "Bob" || name == "Charlie" {
		fmt.Println("Hello, Bob or Charlie!")
	}

}

// I can be wrong, so you can correct me by creating a pull request on my GitHub repository. I will be happy to accept it.
