package main

import "fmt"

func main() {
	// 'x' is a array example that supports five values
	var x [5]int

	// declaring array with filled values
	var y = [5]int{0, 1, 2, 3, 4}

	var a = [4]int{3: 10, 2: 8, 1: 1, 0: 9}

	fmt.Printf("Value of x = %v\n", x)
	fmt.Printf("Value of y = %v\n\n", y)

	fmt.Printf("Value of a = %v\n", a)
}
