package main

import "fmt"

func main() {

	x := 5

	if x > 10 {
		fmt.Println("X is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	fmt.Printf("x is %d\n", x) // prints 5 not 9 as 9 is declared inside if
}
