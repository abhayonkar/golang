package main

import "fmt"

// globle  scope

var x int = 10

func main() {

	x := 5

	if x < 10 {
		fmt.Println("X is", x)

	}

	if x := 9; x > 11 {
		fmt.Println("X is greater than 11")

	} else {
		fmt.Println("X is less than 11")
	}

	fmt.Println(x)
	fmt.Println("X doesn't print 10 as the local varible overrides global")
}
