# The -gcflags "-N -l" is necessary to prevent the Go compiler from inlining to
# Go no-op function. See:
# https://code.google.com/p/go/issues/detail?id=3363
# https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/6m8bZYik8Ss

go test -gcflags "-N -l" -bench=.

# Results on my machine:
#
# testing: warning: no tests to run
# PASS
# BenchmarkGoToCNoOp-8    10000000               155 ns/op
# BenchmarkGoOnlyNoOp-8   500000000                3.15 ns/op
# ok      github.com/draffensperger/go-interlang/benchmarks/go_to_c_no_op 3.614s
