# Calling C++ from Go using SWIG

Many Unix systems include SWIG, for On Mac OS X, run `brew install swig`.

To run the example, simply run `go build` which takes care of calling SWIG, and
then run the program with `./swig`.

## Diving into the auto-generated wrapper code

Run `go build -work` to save the temporary build folder and print out it's
output. Once in the work folder, the interesting auto-generated files are found
via
`cd github.com/draffensperger/go-interlang/go_to_cxx/swig/point/_obj`.

The `point_wrap.cxx` file includes an `extern "C"` section where it makes
wrapper functions for instantiating a `Point` class, and calling the
`distance_to` method on an instance. It also generates a fair bit of other
boilerplate code.

The `point.go` file provides Go functions that call into those `extern "C"`
wrapper functions to provide the Go interface for the Point class.

Together those files provide the glue for Go to easily call the C++ methods on
the Point class in a Go-friendlier fashion.

## Helpful links

SWIG documentation: [swig.org](http://swig.org/)

`go` command documentation (especially "Calling between Go and C")[https://golang.org/cmd/go/]

SWIG Go examples: (github.com/swig/swig/tree/master/Examples/go)[https://github.com/swig/swig/tree/master/Examples/go]

Zac Gross's Go to C++ SWIG example at 
[github.com/zacg/simplelib](https://github.com/zacg/simplelib), and related blog
post
["Calling C++ code form Go with SWIG"](http://zacg.github.io/blog/2013/06/06/calling-c-plus-plus-code-from-go-with-swig/)
.
