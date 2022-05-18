package main

import "fmt"

func main() {
	n, _ := fmt.Println("hello")
	fmt.Println(n)
	//fmt.Println(_err)

	//foo()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
}

func foo() {
	fmt.Println("I'm in foo.")
}
