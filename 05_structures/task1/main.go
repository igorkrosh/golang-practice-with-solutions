package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var person Person
	person = Person{"Alex", 21}

	fmt.Println("Name: ", person.Name)
	fmt.Println("Age: ", person.Age)
}
