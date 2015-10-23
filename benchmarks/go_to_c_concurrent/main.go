package main

// #include "harmonic_range.h"
import "C"
import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

var partialSums []C.double
var wg sync.WaitGroup

func calcPartialSum(goroutine, start, end int) {
	defer wg.Done()
	partialSums[goroutine] = C.harmonic_range(C.int(start), C.int(end))
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	goroutines, err := strconv.Atoi(os.Args[2])
	if err != nil {
		goroutines = runtime.NumCPU()
	}
	fmt.Printf("\n\nGo wth %v goroutines calling C and GOMAXPROCS=%v:\n", goroutines,
		os.Getenv("GOMAXPROCS"))
	fmt.Printf("Sum for n=1..%v of 1/n =\n", n)

	wg.Add(goroutines)
	partialSums = make([]C.double, goroutines)
	rangeLen := n / goroutines
	for i := 0; i < goroutines; i++ {
		start := i*rangeLen + 1
		end := start + rangeLen - 1
		go calcPartialSum(i, start, end)
	}
	wg.Wait()
	sum := C.double(0.0)
	for i := 0; i < goroutines; i++ {
		sum += partialSums[i]
	}
	fmt.Printf("%v", sum)
}
