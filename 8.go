package main

import "fmt"

func main() {
	fmt.Println(swap(5, 4))
}

func swap(i, j int) (int, int) {
	return j, i
}
