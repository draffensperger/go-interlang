# Calling a Go dynamically linked library from C++

To make an run this example, run `./make_and_run.sh`.

The key tricks here are setting the include and link parameters for `g++`, using
`-buildmode=c-shared` for `go build` (see `Makefile`) and then setting the 
library path environment variable for the compiled program to find the library 
(see `make_and_run.sh`).
