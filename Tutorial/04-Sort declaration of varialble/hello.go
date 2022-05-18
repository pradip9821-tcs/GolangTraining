package main

import "fmt"

func main() {
	// Declaration of variable
	a := 10
	fmt.Println(a)

	// Declaration of function
	b, err := fmt.Println(a)
	fmt.Println(b, " ", err)

	// Function return two value and we want one.

	c, _ := fmt.Println(a)
	fmt.Println(c)
}
