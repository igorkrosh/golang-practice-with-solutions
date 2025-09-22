package main

import "fmt"

func main() {
	obj := make(map[string]int)

	obj["Bob"] = 20
	obj["John"] = 25
	obj["Amy"] = 18

	for key, value := range obj {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}
