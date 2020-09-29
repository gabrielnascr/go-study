package main

import "fmt"

// Simple return
func sum(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

// Function that takes a function as a parameter
func executeFunction(f func(string) string, v string) {
	aux := f(v)
	fmt.Println(aux)
}

// Passing infinite numbers
func vIntegers(i ...int) {
	fmt.Println(i)
}

func main() {
	// helloWorld()

	fmt.Println(sum(1, 2))
	fmt.Println(sub(4, 1))

	hello := func(s string) string {
		return "Hello, " + s
	}

	executeFunction(hello, "World")
	vIntegers(1, 2, 3, 5, 4, 7, 8, 9, 7, 4, 4)
}
