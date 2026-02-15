package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello user, What is your name ?")

	// 1. Using Scanf - where the input is taken till space

	// var name string
	// fmt.Scanf("%s", &name)
	// fmt.Printf("Hello, %s!\n", name)

	// 2. Using ReadString - where the input is taken till new line

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("Hello, " + name + "!")

}
