package main

import "fmt"

// When using //export, the "C" preamble can only have,
// declarations (no definitions), because it gets copied twice.

// #include "adder.h"
// #include "callback.h"
import "C"

var total int

// The //export allows C to call the Go func
//export GoTotalCallback
func GoTotalCallback(callbackTotal int) {
	fmt.Printf("Go callback got total %d\n", callbackTotal)
	total = callbackTotal
}

func main() {
	fmt.Printf("Go says: about to call C callback add\n")
	// With cgo you can't call C function pointers directly,
	// but you can pass then to C functions that can call them
	C.add(40, 2, C.callback_fn(C.c_to_go_callback))
	// C.add(40, 2, unsafe.Pointer(&C.c_to_go_callback))
	fmt.Printf("Go says: the result was %d\n", total)
}
