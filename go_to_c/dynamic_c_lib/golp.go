package main

/*
#cgo linux CFLAGS: -I./lpsolve
#cgo linux LDFLAGS: -L./lpsolve -llpsolve55 -Wl,-rpath=./lpsolve
#cgo darwin CFLAGS: -I/opt/local/include/lpsolve
#cgo darwin LDFLAGS: -L/opt/local/lib -llpsolve55
#include <lp_lib.h>
*/
import "C"

import "runtime"

type LP struct {
	ptr *C.lprec
}

func NewLP(rows, cols int) *LP {
	l := new(LP)
	l.ptr = C.make_lp(C.int(rows), C.int(cols))
	runtime.SetFinalizer(l, deleteLP)
	l.SetAddRowMode(true)
	return l
}

func deleteLP(l *LP) {
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
	cRow := make([]C.double, len(row)+1)
	cRow[0] = 0.0
	for i := 0; i < len(row); i++ {
		cRow[i+1] = C.double(row[i])
	}
	C.add_constraint(l.ptr, &cRow[0], C.int(ct), C.double(rightHand))
}

func (l *LP) SetObjFn(row []float64) {
	cRow := make([]C.double, len(row)+1)
	cRow[0] = 0.0
	for i := 0; i < len(row); i++ {
		cRow[i+1] = C.double(row[i])
	}
	C.set_obj_fn(l.ptr, &cRow[0])
}

type SolutionType int

const (
	NOMEMORY   = -2
	OPTIMAL    = iota // OPTIMAL == 0
	SUBOPTIMAL        // SUBOPTIMAL == 1
	INFEASIBLE        // INFEASIBLE == 2
	UNBOUNDED
	DEGENERATE
)

func (l *LP) Solve() SolutionType {
	return SolutionType(C.solve(l.ptr))
}

func (l *LP) Objective() float64 {
	return float64(C.get_objective(l.ptr))
}

func (l *LP) Variables() []float64 {
	numCols := int(C.get_Ncolumns(l.ptr))
	cRow := make([]C.double, numCols)
	C.get_variables(l.ptr, &cRow[0])
	row := make([]float64, numCols)
	for i := 0; i < numCols; i++ {
		row[i] = float64(cRow[i])
	}
	return row
}
