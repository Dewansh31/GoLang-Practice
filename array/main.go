package main

import (
	"fmt"
	"mylearning/myutils"
)

func main() {
	myutils.PrintMessage("Learning Arrays in Go!")

	var names [3]string
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	fmt.Println("Array contents:", names)

}
