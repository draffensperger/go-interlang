# Go inter-language call benchmarks

These benchmarks explore the cost of calls from Go to C and Ruby to Go, as well
as HTTP calls and Go to C call concurrency implications.

To run all the benchmarks, run `./make_and_run.sh`, which compiles the code and
calls the shell scripts for the various benchmarks.

For a more general comparison of language performance, see 
[The Computer Language Benchmarks Game](http://benchmarksgame.alioth.debian.org/)
(source copied on GitHub at 
[github.com/nbraud/benchmarksgame](https://github.com/nbraud/benchmarksgame)).

The raw results from runs of the benchmarks on my machine in
[results.md](https://github.com/draffensperger/go-interlang/blob/master/benchmarks/results.md).

Two helpful links on the performance and concurrency implications of Go to C
calls are a golang-nuts thread
[What is the overhead of calling a C function from Go?](https://groups.google.com/forum/#!topic/golang-nuts/RTtMsgZi88Q) and a Stack Overflow question
[C code and goroutine scheduling](http://stackoverflow.com/questions/28354141/c-code-and-goroutine-scheduling).

## Go to C call cost

The most direct way to measure the cost of a Go to C call is via a no-op
benchmark of a cgo call and a pure Go call. That benchmark can be found in the
`go_to_c_no_op` folder. The benchmark depends on disabling function inlining in
the Go compiler via the `-gcflags "-N -l"` flag.

The no-op benchmark on my machine found the cost of a Go to C call to be about
160 ns, and the cost of a Go call to be about 3ns, making a cgo call about 50x
slower than a Go call.

That overhead broadly concurs with the 
[golang-nuts thread on cgo overhead](https://groups.google.com/forum/#!topic/golang-nuts/RTtMsgZi88Q)
in that the difference is noticable due to the need for a stack change, though some 
people there seemed to have a ratio closer to 9-25x. Someone there did have the 
same value of 160ns for a cgo call. The thread is from 2013 so the Go runtime
may have changed since their benchmarks were run.

A less direct measure is comparing how a Go to C calculation of the harmonic
series (with a cgo call for each item of the series) compares with the pure Go
and pure C versions. The Go to C harmonic seris with n-times cgo calls took
about 40x longer than the pure C version, which broadly fits with the ratio of
50x for the no-op benchmark since its time is dominated by the cost of the cgo
calls.

Basically calling C from Go incurs a cost but you can still do millions /
second (1s / 160ns = 6.25 million), so it makes sense to use cgo calls, just not
in a tight, high performance loop.

## Ruby to Go call cost

The Ruby FFI cost from the no-op Ruby to Go vs. pure Ruby function call
benchmark got a ratio of about 28x.

The Ruby to Go harmonic series calculation with n-times FFI calls took 24x as
long as the pure Ruby version, which fits well with the no-op benchmark assuming
that Go code running in those FFI calls gave a slight speedup compared to pure
Ruby + pure FFI cost.

The Go calculation for the harmonic series (0.55s)was about 10x faster than the 
pure Ruby one (5.5s). Ruby using a single FFI call (0.66s) was 9x faster than
the pure Ruby version.

So calling Go from Ruby can be helpful, but again, avoid repeated calls from
Ruby in a tight loop.

## Concurrency implications

As mentioned in the 

## HTTP call cost


