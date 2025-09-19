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

	switch number {
	case 1:
		fmt.Println("Weekday")
	case 2:
		fmt.Println("Weekday")
	case 3:
		fmt.Println("Weekday")
	case 4:
		fmt.Println("Weekday")
	case 5:
		fmt.Println("Weekday")
	case 6:
		fmt.Println("Weekend")
	case 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Incorrect number")
	}

}
