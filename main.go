package main

import (
	"fmt"
	lexer "github.com/Shavvey/lisper/lexer"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Printf("The value is %d\n", foo(3))
}

func foo(a int) int {
	return a
}
