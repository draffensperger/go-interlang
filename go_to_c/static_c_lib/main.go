package main

import "fmt"
import "unsafe"

// #cgo CFLAGS: -I${SRCDIR}/greetings
// #cgo LDFLAGS: ${SRCDIR}/greetings/greetings.a
// #include <stdlib.h>
// #include <greetings.h>
import "C"

func main() {
	fmt.Printf("Hi from Go: get ready for multi-lingual greetings!\n")

	john := C.CString("John")
	johannes := C.CString("Johannes")
	C.greet_in_english(john)
	C.greet_in_german(johannes)

	fmt.Printf("Bye from Go: freeing memory allocated on C heap...")
	C.free(unsafe.Pointer(john))
	C.free(unsafe.Pointer(johannes))
}
