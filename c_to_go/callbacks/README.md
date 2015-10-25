# Calling from Go to C and back

To run this example run `go build` then `./to_c_and_back`.

Note: When using //export, the "C" preamble can only have, declarations 
(no definitions), because it gets copied twice. (See 
[cgo docs](https://golang.org/cmd/cgo/)).
