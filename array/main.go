package main

import (
	"fmt"
)

func main() {
	// myutils.PrintMessage("Learning Arrays in Go!")

	// var names [3]string
	// names[0] = "Alice"
	// names[1] = "Bob"
	// names[2] = "Charlie"

	// fmt.Println("Array contents:", names)

	// var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Numbers array:", numbers, "Length:", len(numbers))

	var price [5]int
	fmt.Println("Price array before assignment:", price)

	/*

		The text explains how elements in a Go array or slice are initialized to their zero values by default. The zero value depends on the variable's type.
		- Numeric types (int, float, etc.) have a zero value of 0.
		- Strings have a zero value of an empty string ("").
		- Boolean types have a zero value of false.
		- Pointers or complex types have a zero value of nil.

	*/

}
