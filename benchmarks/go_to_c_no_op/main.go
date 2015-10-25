package main

// #include "no_op.h"
import "C"

func GoToCNoOp() {
	C.no_op()
}

//export GoOnlyNoOp
func GoOnlyNoOp() {
}

func main() {
}
