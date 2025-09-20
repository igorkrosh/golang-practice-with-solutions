package main

import "fmt"

func main() {

	fmt.Println("Factorial 5: ", factorial(5))
}

func factorial(n int) (result int) {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}
