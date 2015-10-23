package main

import "C"
import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func AddHarmonic(totalSoFar C.double, i C.int) C.double {
	return totalSoFar + 1.0/C.double(i)
}

func HarmonicRange(start, end C.int) C.double {
	sum := C.double(0.0)
	for i := C.int(start); i <= end; i++ {
		sum = AddHarmonic(sum, i)
	}
	return sum
}

var cpuSums []C.double
var wg sync.WaitGroup

func calcRangeSum(cpu, i, j C.int) {
	defer wg.Done()
	cpuSums[cpu] = HarmonicRange(i, j)
}

func main() {
	n_int, _ := strconv.Atoi(os.Args[1])
	n := C.int(n_int)
	procs, err := strconv.Atoi(os.Getenv("GOMAXPROCS"))
	if err != nil {
		procs = runtime.NumCPU()
	}
	numCpu := C.int(procs)
	fmt.Printf("\n\nGo concurrent GOMAXPROCS=%v:\n", os.Getenv("GOMAXPROCS"))
	fmt.Printf("Sum for n=1..%v of 1/n =\n", n)

	wg.Add(int(numCpu))
	cpuSums = make([]C.double, numCpu)
	rangeLen := n / numCpu
	for cpu := C.int(0); cpu < numCpu; cpu++ {
		start := cpu*rangeLen + 1
		end := start + rangeLen - 1
		fmt.Printf("Range %v is %v to %v\n", cpu, start, end)
		go calcRangeSum(cpu, start, end)
	}
	wg.Wait()
	sum := C.double(0.0)
	for cpu := C.int(0); cpu < numCpu; cpu++ {
		sum += cpuSums[cpu]
	}
	fmt.Printf("%v", sum)
}
