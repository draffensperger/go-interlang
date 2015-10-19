# Calling from Go to C and back

To run this example run `go build` then `./to_c_and_back`.

This also shows how to get around cgo's inability to call C function pointers:
by passing the C function pointer to a C function that then calls it.

This is a great writeup of C/Go callbacks using a very similar adder example:
[Callbacks with cgo](http://cheesesun.blogspot.com/2010/04/callbacks-in-cgo.html).

Here's a StackOverflow question with a similar example:
[Call go functions from C](http://stackoverflow.com/questions/6125683/call-go-functions-from-c)

The repo [fiorix/gocp](https://github.com/fiorix/gocp) from is an example of
exposing Go concurrency primitives as a C library. It uses a similar trick to
call a C function pointer via a bridge function `bridge_void_func`, 
see its [gocp.c](https://github.com/fiorix/gocp/blob/master/src/gocp/gocp.c)

Note: When using //export, the "C" preamble can only have, declarations 
(no definitions), because it gets copied twice. (See 
[cgo docs](https://golang.org/cmd/cgo/)).
