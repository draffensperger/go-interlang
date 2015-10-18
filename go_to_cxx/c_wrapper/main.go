package main

// #include "wrap_point.hpp"
import "C"

import "fmt"

func main() {
	fmt.Println("Hi from Go, about to calculate distance in C++ ...")
	distance := C.distance_between(1.0, 1.0, 2.0, 2.0)
	fmt.Printf("Go has result, distance is: %v\n", distance)
}
