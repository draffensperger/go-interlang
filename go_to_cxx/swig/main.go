package main

import (
	"fmt"

	"github.com/draffensperger/go-interlang/go_to_cxx/swig/point"
)

func main() {
	fmt.Println("Hi from Go, SWIG did the wrapping so we can calculate distance...")
	p1 := point.NewPoint(1.0, 1.0)
	p2 := point.NewPoint(2.0, 2.0)
	distance := p1.Distance_to(p2)
	fmt.Printf("Go has result, distance is: %v\n", distance)
}
