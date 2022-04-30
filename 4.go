package main

//import "fmt"
//import "math/rand"
//import "strings"

import (
	"fmt"
	"math/rand"
)

func main() {
	var i = (rand.Intn(10))

	if i < 10 {
		fmt.Println(rand.Int())
	}
}
