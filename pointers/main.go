package main

func main() {

	num := 42
	numPtr := &num

	println("Value of num:", num)
	println("Address of num:", &num)
	println("Value of numPtr:", numPtr)
	println("Value pointed to by numPtr:", *numPtr)

	*numPtr = 100
	println("Updated value of num:", num)
}
