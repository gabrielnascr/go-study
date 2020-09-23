package main

import (
	"fmt"
)

var ( // múltiples variables
	name  = "Gabriel"
	idade = 15
)

func main() {
	// Declaring multi variables on the same line
	var x, y int = 5, 4

	var a string
	var b string

	a = "First Text"
	b = "Second Text"

	fmt.Printf("Value of a = %s\n", a)
	fmt.Printf("Value of b = %s\n\n", b)

	fmt.Printf("Value of x = %d\n", x)
	fmt.Printf("Value of y = %d\n\n", y)

}
