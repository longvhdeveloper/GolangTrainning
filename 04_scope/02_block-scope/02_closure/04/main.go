package main

import "fmt"

func wrapper() func() int {
	x := 40
	return func() {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
