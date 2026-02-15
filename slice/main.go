package main

import "fmt"

func main() {

	// https://golangforall.com/en/post/golang-slice.html

	// var numbers = []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 6, 7, 8, 9, 10)

	numbers := make([]int, 5, 10)
	fmt.Println("Numbers slice:", numbers, "\nLength:", len(numbers), "\nCapacity:", cap(numbers), "\nData type:", fmt.Sprintf("%T", numbers))

}
