# Call C++ from Go with cgo via a C wrapper

To run the example, run `go build` then `./c_wrapper`.

## Exploring C++ name manging with nm utility

The [nm](http://linux.die.net/man/1/nm) utility explores symbol names in object 
files.

To see what the C++ name mangling looks like, run `g++ -c point.cxx` then 
`nm point.o | grep distance`. You will see how C++ exports the `distance_to`
function to include some extra decoration to show e.g. how it's connected to the
`Point` class.

To see the cleaner `extern "C"` name, run `g++ -c *.cxx`, and then `nm
wrap_point.o | grep distance`. You'll see an external reference to the C++
mangled name for `Point::distance_to`, but you'll also see the cleaner
`_distance_between` name which Go can call.

## Other similar examples

Here are a couple other similar examples of calling C++ from Go without SWIG
using a C wrapper:

[github.com/wangkuiyi/go-cpp/](https://github.com/wangkuiyi/go-cpp/)

[github.com/burke/howto-go-with-cpp](https://github.com/burke/howto-go-with-cpp)
