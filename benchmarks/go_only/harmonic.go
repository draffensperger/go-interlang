package main

import "C"
import (
	"fmt"
	"os"
	"strconv"
)

//export AddHarmonic
func AddHarmonic(totalSoFar C.double, i C.int) C.double {
	return totalSoFar + 1.0/C.double(i)
}

//export HarmonicSum
func HarmonicSum(n C.int) C.double {
	sum := C.double(0.0)
	for i := C.int(1); i <= n; i++ {
		sum = AddHarmonic(sum, i)
	}
	return sum
}

func main() {
	n_int, _ := strconv.Atoi(os.Args[1])
	n := C.int(n_int)
	fmt.Printf("\n\nGo only:\nSum for n=1..%v of 1/n =\n", n)
	fmt.Printf("%v", HarmonicSum(n))
}
