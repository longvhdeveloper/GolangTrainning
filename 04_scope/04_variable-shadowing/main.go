package main

import "fmt"

func max(x int) {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max now is now the result , not the function
}
 // don't do this; bad coding pratice to shadow variable
