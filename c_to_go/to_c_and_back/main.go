package main

import "fmt"

// #include "adder.h"
// #include "callback.h"
import "C"

var total int

// The //export allows C to call the Go func

//export GoTotalCallback
func GoTotalCallback(callbackTotal C.int) {
	fmt.Printf("Go callback got total %d\n", callbackTotal)
	total = int(callbackTotal)
}

func main() {
	fmt.Printf("Go says: about to call C callback add\n")

	// With cgo you can't call C function pointers directly,
	// but you can pass then to C functions that can call them
	C.add(40, 2, C.callback_fn(C.c_to_go_callback))
	fmt.Printf("Go says: 1st result is %d\n\n", total)

	C.add_with_go_callback(100, 1)
	fmt.Printf("Go says: 2nd result is %d\n", total)
}
