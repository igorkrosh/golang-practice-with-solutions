package main

import "fmt"

func main() {
	pi := 3.14159

	var radius int

	fmt.Print("Enter the radius: ")
	_, err := fmt.Scan(&radius)

	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}

	s := pi * float64(radius*radius)

	fmt.Println("Area of a circle:", s)
}
