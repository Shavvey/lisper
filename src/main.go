package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	fmt.Printf("The value is %d\n", foo(3))
}

func foo(a int) int {
	return a
}
