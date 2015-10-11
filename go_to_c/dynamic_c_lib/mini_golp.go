// Simplified excerpts from golp package for calling LPSolve
// Full package at: https://github.com/draffensperger/golp
// LPSolve C library documentation: http://lpsolve.sourceforge.net/5.5/
package main

/*
// The "linux" and "darwin" below are platform-based build constraints
// See: https://golang.org/pkg/go/build/

// On Linux, assume that lpsolve is extracted to a subdirectory, e.g. via:
// LP_URL=http://sourceforge.net/projects/lpsolve/files/lpsolve/5.5.2.0/lp_solve_5.5.2.0_dev_ux64.tar.gz
// mkdir -p lpsolve
// wget -qO- $LP_URL | tar xvz -C lpsolve
#cgo linux CFLAGS: -I./lpsolve
#cgo linux LDFLAGS: -L./lpsolve -llpsolve55 -Wl,-rpath=./lpsolve

// On Mac, assume LPSolve is installed via MacPorts with "port install lpsolve"
#cgo darwin CFLAGS: -I/opt/local/include/lpsolve
#cgo darwin LDFLAGS: -L/opt/local/lib -llpsolve55

#include <lp_lib.h>
*/
import "C"

import (
	"fmt"
	"runtime"
)

type LP struct {
	ptr *C.lprec
}

func NewLP(rows, cols int) *LP {
	l := new(LP)
	l.ptr = C.make_lp(C.int(rows), C.int(cols))
	runtime.SetFinalizer(l, deleteLP)
	return l
}

func deleteLP(l *LP) {
	fmt.Println("Go finalizer says: deleting C linear program structure")
	C.delete_lp(l.ptr)
}

type ConstraintType int

const (
	_  = iota // don't use 0
	LE        // LE == 1
	GE        // GE == 2
	EQ        // EQ == 3
)

func (l *LP) AddConstraint(row []float64, ct ConstraintType, rightHand float64) {
	C.add_constraint(l.ptr, (*C.double)(&row[0]), C.int(ct), C.double(rightHand))
}

func (l *LP) SetObjFn(row []float64) {
	C.set_obj_fn(l.ptr, (*C.double)(&row[0]))
}

func (l *LP) SetMaximize() {
	C.set_maxim(l.ptr)
}

func (l *LP) Solve() {
	C.solve(l.ptr)
}

func (l *LP) Objective() float64 {
	return float64(C.get_objective(l.ptr))
}

func (l *LP) Variables() []float64 {
	numCols := int(C.get_Ncolumns(l.ptr))
	row := make([]float64, numCols)
	C.get_variables(l.ptr, (*C.double)(&row[0]))
	return row
}
