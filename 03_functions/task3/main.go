package main

import "fmt"

func main() {
	str1 := "foo"
	str2 := "bar"

	fmt.Println("Str1 before swap: ", str1)
	fmt.Println("Str2 before swap: ", str2)

	swap(&str1, &str2)

	fmt.Println("Str1 after swap:", str1)
	fmt.Println("Str2 after swap:", str2)

}

func swap(str1 *string, str2 *string) {
	temp := *str1

	*str1 = *str2
	*str2 = temp
}
