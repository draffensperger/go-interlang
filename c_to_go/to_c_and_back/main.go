package main

import "fmt"

// #include "adder.h"
import "C"

// The //export allows C to call the Go func
//export GiveGoTotal
func GiveGoTotal(total C.int) {
	fmt.Printf("Go says: got total from C %d\n", total)
}

func main() {
	fmt.Printf("Go says: calling C to add numbers..\n")
	C.add_and_give_go_total(40, 2)
}
