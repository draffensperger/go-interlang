package main

import (
	"fmt"
	"unsafe"
)

/*
// This preamble is interpreted as C code
// See cgo docs: https://golang.org/cmd/cgo/
#include <stdio.h>
#include <stdlib.h>
int add(int x, int y) {
	printf("C says: adding %i + %i\n", x, y);
	return x + y;
}
*/
import "C"

func main() {
	total := C.add(10, 20)
	fmt.Printf("Go says: result is %v\n", total)

	gostr := "Go can call C's puts from stdio"
	// C.CString uses C heap, we must free it
	var cstr *C.char = C.CString(gostr)
	// defer is convenient for clean-up
	defer C.free(unsafe.Pointer(cstr))
	C.puts(cstr)
}
