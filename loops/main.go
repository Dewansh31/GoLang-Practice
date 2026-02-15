package main

func main() {

	// for i := 1; i <= 5; i++ {
	// 	println("Iteration:", i)
	// }

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}

	name := "GoLang"
	for _, char := range name {
		println("Character:", string(char))
	}

	/*

		In the Go programming language for iterating over collections.
		The range keyword simplifies iteration over slices, arrays, maps, and strings.
		It provides both the index (or key) and the value of each element during iteration.
		In the example provided, range numbers iterates through a slice, setting the index and value on each loop.

	*/

}
