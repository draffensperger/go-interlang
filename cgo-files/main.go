package main

import "fmt"

// #include "adder.h"
import "C"

func main() {
	z := C.adder(10, 20)
	fmt.Printf("Hello from Go: the result was: %v\n", z)
}
