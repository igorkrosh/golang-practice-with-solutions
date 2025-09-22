package main

import "fmt"

func main() {
	var slice []string

	slice = append(slice, "aaa")
	slice = append(slice, "bbb")
	slice = append(slice, "ccc")

	fmt.Println("Slice before: ", slice)

	slice = append(slice[:1], slice[2:]...)

	fmt.Println("Slice after: ", slice)
}
