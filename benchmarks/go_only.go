package main

import (
	"fmt"
	"os"
	"strconv"
)

// calcualte sum of 1 .. n = 1 / n
func main() {
	var n int
	var err error
	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Must enter int for N for sum of 1 .. N")
		}
	} else {
		n = 100
	}

	var total float64 = 0.0
	for i := 1; i < n; i++ {
		total += 1.0 / float64(i)
	}

	fmt.Printf("Sum of n = 1..%v of 1/n : %v\n", n, total)
}
