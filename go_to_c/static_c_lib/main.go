package main

import "fmt"
import "unsafe"

// #cgo CFLAGS: -I${SRCDIR}/greetings
// #cgo LDFLAGS: ${SRCDIR}/greetings.a
// #include <stdlib.h>
// #include <greetings.h>
import "C"

func main() {
	fmt.Printf("Go says: static C library greetings coming..\n")

	john := C.CString("John")
	defer C.free(unsafe.Pointer(john))
	johannes := C.CString("Johannes")
	defer C.free(unsafe.Pointer(johannes))

	C.greet_in_english(john)
	C.greet_in_german(johannes)

	fmt.Printf("Go says bye!\n")
}
