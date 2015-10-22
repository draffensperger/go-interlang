package main

import (
	"fmt"
	"os"
	"strconv"
)

// calcualte sum of 1 .. n = n(n + 1) / 2
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

	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}

	fmt.Printf("Sum of 1 .. %v = %v\n", n, total)
}
