package main

import "fmt"

import "C"

//export GoAdder
func GoAdder(x, y int) int {
	fmt.Printf("Go says: adding %v and %v\n", x, y)
	return x + y
}

func main() {} // Required but ignored
