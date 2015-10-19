# Calling from Go to C and back

This also shows how to get around cgo's inability to call C function pointers:
by passing the C function pointer to a C function that then calls it.

Here's a StackOverflow question with a similar example:
[Call go functions from C](http://stackoverflow.com/questions/6125683/call-go-functions-from-c)

The repo [fiorix/gocp](https://github.com/fiorix/gocp) from is an example of
exposing Go concurrency primitives as a C library. It uses a similar trick to
call a C function pointer via a bridge function `bridge_void_func`, 
see its [gocp.c](https://github.com/fiorix/gocp/blob/master/src/gocp/gocp.c)
