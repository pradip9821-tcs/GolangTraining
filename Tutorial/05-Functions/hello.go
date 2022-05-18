package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	fmt.Println(a + b)
	foo()
	var ans = sum(a, b)
	fmt.Println(ans)
}

func foo() {
	fmt.Println("I'm in foo.")
}

func sum(a int, b int) int {
	return a + b
}
