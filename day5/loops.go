package main

import "fmt"

// No need to learn alot of loops ;)
// Go only knows for loop !

// I am a Python developer, learnt two kind of loops: While and for loop
// But here in Go it is only for which performs all the job

func main() {

	// for loop with just a condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop like low level languages. Initialization; Condition; Iterator

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Performs n times any task
	for i := range 3 {
		fmt.Println("range", i)
	}

	// Infinite loop

	for {
		fmt.Println("loop")
		break
	}


	// You can also use continue keyword to skip the current iteration
		
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

}
