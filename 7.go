package main

import (
	"fmt"
)

func main() {
	subt := sub(65, 98)
	fmt.Println(subt)
	a, b, c, d := multiples(7, 8, 9, 10)
	fmt.Println(a, b, c, d)
}

func sub(i, j int) int {
	return j - i
}

func multiples(i, j, k, l int) (int, int, int, int) {
	i = i * 2
	j = j * 3
	k = k + 2
	l = l * 6
	return i, j, k, l

}
