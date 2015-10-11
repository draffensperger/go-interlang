package main

import "fmt"

/*
#include <stdio.h>
int adder(int x, int y) {
	printf("Hello from C: adding %i + %i\n", x, y);
	return x + y;
}
*/
import "C"

func main() {
	z := C.adder(10, 20)
	fmt.Printf("Hello from Go: the result was: %v\n", z)
}
