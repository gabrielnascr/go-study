package main

import (
	"fmt"
)

func main() {
	age := 15

	if age > 18 {
		fmt.Println("Age of majority")
	} else {
		fmt.Println("Minority")
	}
}
