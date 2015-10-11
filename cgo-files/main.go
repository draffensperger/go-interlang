package main

import "fmt"
import "unsafe"

// #include <stdlib.h>
// #include "concat.h"
import "C"

func main() {
	s1 := C.CString("malloc allocated memory (or from CString)")
	s2 := C.CString("must be freed!")
	result := C.concat(s1, s2)
	fmt.Printf("Hello from Go: the result was: \"%v\"\n", C.GoString(result))
	C.free(unsafe.Pointer(result))
	C.free(unsafe.Pointer(s1))
	C.free(unsafe.Pointer(s2))
}
