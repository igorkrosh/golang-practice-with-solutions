package main

import "fmt"

func main() {
	add(1, 2)
	add(10, -2)
	add(43, 0)
}

func add(a int, b int) (result int) {
	fmt.Println("First value: ", a)
	fmt.Println("Secons value: ", b)

	result = a + b

	fmt.Println("Result: ", result)

	return result
}
