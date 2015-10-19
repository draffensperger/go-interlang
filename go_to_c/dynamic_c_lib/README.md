# Calling a dynamically-linked C library from Go

To run this example, you first need install
[LPSolve](http://lpsolve.sourceforge.net/5.5/). See below for instructions.

Then just run `go build` to build the Go executable, and run it with
`./dynamic_c_lib`.

## How to install LPSolve

Here's how you could download and extract the LPSolve library for 64-bit Linux:

```
LP_URL=http://sourceforge.net/projects/lpsolve/files/lpsolve/5.5.2.0/lp_solve_5.5.2.0_dev_ux64.tar.gz
mkdir lpsolve
curl -L $LP_URL | tar xvz -C lpsolve
```

To install LPSolve on Mac OS X, install [MacPorts](https://www.macports.org/),
then run `sudo port install lp_solve`.

## More complete LPSolve Go wrapper

The `lp.go` file is a simplified version of the same file in the
[golp](https://github.com/draffensperger/golp)
package I made to wrap the
[LPSolve](http://lpsolve.sourceforge.net/5.5/)
Mixed Integer Linear Programming (MILP) solver.

For more discussion of concepts in this example, see my blog post 
[Calling a Linear Solver C Library From Go](http://davidraff.com/calling-a-linear-programming-solver-from-go).

## Other links and notes

[golang-nuts: Passing a slice/array pointer between Go/C](https://groups.google.com/forum/#!topic/golang-nuts/DXCIP6pMMN0)

Something I learned by forgetting a \n in printf: 
[Go doesn't flush C streams on program exit](https://github.com/golang/go/issues/8808),
and
[printf doesn't flush unless there is a \n in the format](http://stackoverflow.com/questions/1716296/why-does-printf-not-flush-after-the-call-unless-a-newline-is-in-the-format-string),
so if you try a printf without \n in your C code called by Go the message won't
show up.
