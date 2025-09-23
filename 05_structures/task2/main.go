package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person *Person) greet() {
	fmt.Println("Hello, I'm", person.Name)
}

func main() {
	var person Person
	person = Person{"Alex", 21}

	person.greet()
}
