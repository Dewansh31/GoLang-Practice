package main

import (
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divide by zero.")
	}
	return a / b, nil
}

func main() {

	fmt.Println("Currently learning error handling...")

	ans, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Answer :", ans)
	}

}
