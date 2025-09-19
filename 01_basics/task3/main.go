package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	_, err1 := fmt.Scan(&a)

	if err1 != nil {
		fmt.Println("Input error: ", err1)
		return
	}

	fmt.Print("Enter second number: ")
	_, err2 := fmt.Scan(&b)

	if err2 != nil {
		fmt.Println("Input error: ", err2)
		return
	}

	sum := a + b
	diff := a - b
	mul := a * b
	div := 0.0

	fmt.Println("Sum: ", sum)
	fmt.Println("Difference: ", diff)
	fmt.Println("Multiplication: ", mul)

	if b == 0 {
		fmt.Println("Division by zero is not possible")
	} else {
		div = float64(a) / float64(b)
		fmt.Println("Division: ", div)
	}
}
