package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// No access to x
	// This does not compile
	fmt.Println(x)
}
