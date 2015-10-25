# Go inter-language call benchmarks

To run all the benchmarks, run `./make_and_run.sh`, which compiles the code and
calls the shell scripts for the various benchmarks.

For a more general comparison of language performance, see 
[The Computer Language Benchmarks Game](http://benchmarksgame.alioth.debian.org/)
(source copied on GitHub at 
[github.com/nbraud/benchmarksgame](https://github.com/nbraud/benchmarksgame)).

The raw results from runs of the benchmarks on my machine in
[results.md](https://github.com/draffensperger/go-interlang/blob/master/benchmarks/results.md).

http://stackoverflow.com/questions/28354141/c-code-and-goroutine-scheduling

Two great links on the performance and concurrency implications of Go to C 
calls are a golang-nuts thread
[What is the overhead of calling a C function from Go?](https://groups.google.com/forum/#!topic/golang-nuts/RTtMsgZi88Q) and a Stack Overflow question
[C code and goroutine scheduling](http://stackoverflow.com/questions/28354141/c-code-and-goroutine-scheduling).

## Go to C call cost

go to c calls max cost:
(60+56)/800000000
= 1.45e-7
600 cycles

1/4e9 is time for one cycle

145 nanoseconds at the most

A direct benchmark via the `go_to_c_no_op` test got the cost of a Go to C call
as around 143 ns/op which is very similar.

millisecond e-3
microsecond e-6
nanosecond e-9

14.5 e-8
K e3
M e6
G e9
4.0 Ghz

## Ruby to Go call cost

## Concurrency implications

## HTTP call cost
