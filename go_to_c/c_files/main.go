package main

import "fmt"
import "unsafe"

// #include <stdlib.h>
// #include "concat.h"
import "C"

func main() {
	s1 := C.CString("malloc allocated memory (or from CString)")
	defer C.free(unsafe.Pointer(s1))

	s2 := C.CString("must be freed!")
	defer C.free(unsafe.Pointer(s2))

	result := C.concat(s1, s2)
	defer C.free(unsafe.Pointer(result))

	fmt.Printf("Hello from Go: the result was: \"%v\"\n", C.GoString(result))
}
