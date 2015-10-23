package main

// #include "harmonic_sum.h"
import "C"
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n_int, _ := strconv.Atoi(os.Args[1])
	n := C.int(n_int)
	fmt.Println("\n\nGo with 1 time C harmonic_sum call:")
	fmt.Printf("Sum for n=1..%v of 1/n =\n", n)
	fmt.Printf("%v", C.harmonic_sum(n))
}
