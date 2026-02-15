package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	// 1. create a struct type named "Person" with fields "Name" (string) and "Age" (int).
	var person Person
	person.Name = "John"
	person.Age = 30

	fmt.Println("Person : ", person)

	// 2. create an instance of the "Person" struct and assign values to its fields, then print the struct.
	person2 := Person{Name: "Alice", Age: 25}
	fmt.Println("Person 2 : ", person2)

	// 3. new keyword to create a pointer to a struct instance, then assign values to the fields and print the struct through the pointer.
	personPtr := new(Person)
	personPtr.Name = "Bob"
	personPtr.Age = 40
	fmt.Println("Person Pointer : ", *personPtr)

}
