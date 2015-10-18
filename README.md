# Go interlanguage call examples

Code samples for calls between Go and C/C++ and calling Go from dynamic languages.

## Calls from Go to C (`go_to_c` folder)

These examples use [cgo](https://golang.org/cmd/cgo/) to enable calls to C.

- [c_in_comment](https://github.com/draffensperger/go-interlang/blob/master/go_to_c/c_in_comment/main.go): Calling a C snippet in the cgo comment
- Calling C source files in your Go project folder
- Calling a C statically-linked library (a `.a` file)
- Calling a C dynamically-linked library (a `.so` or `.dylib` file)

## Calls from Go to C++ (`go_to_cxx`)

C++ has more complex calling conventions (e.g. function overloading, inheritance, templates) and so it uses [name mangling](https://en.wikipedia.org/wiki/Name_mangling#Name_mangling_in_C.2B.2B) which adds a step when calling it from Go. Below are ways to do it.

- Calling C++ via a C wrapper function
- Calling C++ via an auto-generated [SWIG](http://www.swig.org/) wrapper

## Calls from C/C++ to Go

- Calling from Go to C and back again in your Go project (gets around cgo function pointer limitation)
- Calling a Go static library with `buildmode=c-archive`
- Calling Go from C using [gccgo](https://golang.org/doc/install/gccgo)
- You can also call from C/C++ to a Go dynamically-linked library (same concept as dynamic langs below)

## Calls from Python/Node.js/Ruby/Java to Go

Go now allows building a C-compatible dynamically-linked library with `buildmode=c-shared`. That allows any language that can call C dynamic libraries to call Go.

- Call from Python via [ctypes](https://docs.python.org/2/library/ctypes.html)
- Call from Node.js via [ffi](https://github.com/node-ffi/node-ffi) npm module
- Call from Ruby via [ffi](https://github.com/ffi/ffi) gem
- Call from Java via [JNA](https://github.com/java-native-access/jna)

# License

[MIT Licensed](http://opensource.org/licenses/MIT).
