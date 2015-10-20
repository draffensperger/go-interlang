# Example of using gccgo for Go and C interoperability

This is an example of a C program calling a Go function that gets statically
compiled into the program. To see the example install gccgo (`sudo apt-get install gccgo` on Ubuntu), then run `make` and `./main`.

This example was adapted from a 
[stack overflow answer](http://stackoverflow.com/questions/6125683/call-go-functions-from-c)
.

## Helpful links

Gccgo documentation: [golang.org/doc/install/gccgo](https://golang.org/doc/install/gccgo)

A writeup on using gccgo to link C++ and Go: [Statically Link C++ Code With Go Code](https://cxwangyi.wordpress.com/2011/03/28/statically-linking-c-code-with-go-code/)

Note that as of this writing (October 2015), gccgo doesn't build on Mac OS X.
See 
[github.com/golang/go/issues/463](https://github.com/golang/go/issues/463) and 
[github.com/golang/go/issues/10727](https://github.com/golang/go/issues/10727).
