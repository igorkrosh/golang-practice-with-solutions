package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter number: ")
	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}

	if number%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
