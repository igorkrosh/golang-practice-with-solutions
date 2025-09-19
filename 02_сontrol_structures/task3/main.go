package main

import "fmt"

func main() {
	var n int
	var sum int

	fmt.Print("Enter N: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}

	for n > 0 {
		sum += n
		n--
	}

	fmt.Println("Sum: ", sum)
}
