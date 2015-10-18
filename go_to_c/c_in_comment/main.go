package main

import (
	"fmt"
	"unsafe"
)

/*
// This comment block is interpreted as C code
// See cgo docs: https://golang.org/cmd/cgo/
#include <stdio.h>
#include <stdlib.h>
int add(int x, int y) {
	printf("Hello from C: adding %i + %i\n", x, y);
	return x + y;
}
*/
import "C"

func main() {
	gostr := "Go can call C's puts from #included stdio"
	// CString allocates on the C heap, so we need to free it
	var cstr *C.char = C.CString(gostr)
	// defer makes cleanup easy esp. with multiple return points
	defer C.free(unsafe.Pointer(cstr))
	C.puts(cstr)

	total := C.add(10, 20)
	fmt.Printf("Hello from Go: the result was: %v\n", total)
}
