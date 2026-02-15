package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Currently learning functions...")
	sum := add(3, 5)
	fmt.Printf("The sum of 3 and 5 is: %d\n", sum)
}
