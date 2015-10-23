package main

// #include "add_harmonic.h"
import "C"
import (
	"fmt"
	"os"
	"strconv"
)

func HarmonicSum(n C.int) C.double {
	sum := C.double(0.0)
	for i := C.int(1); i <= n; i++ {
		sum = C.add_harmonic(sum, i)
	}
	return sum
}

func main() {
	n_int, _ := strconv.Atoi(os.Args[1])
	n := C.int(n_int)
	fmt.Println("\n\nGo with n times C add_harmonic calls:")
	fmt.Printf("Sum for n=1..%v of 1/n =\n", n)
	fmt.Printf("%v", HarmonicSum(n))
}
