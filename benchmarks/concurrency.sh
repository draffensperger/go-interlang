N=8000000000

echo
echo =======================================
echo --- Go to C and concurrency ---
echo =======================================

time go_only/main $N
time c_only/main $N

time GOMAXPROCS=1 go_concurrent/main $N 1
time GOMAXPROCS=1 go_to_c_concurrent/main $N 1

time GOMAXPROCS=1 go_concurrent/main $N 8
time GOMAXPROCS=1 go_to_c_concurrent/main $N 8

time GOMAXPROCS=8 go_concurrent/main $N 8
time GOMAXPROCS=8 go_to_c_concurrent/main $N 8

time GOMAXPROCS=4 go_concurrent/main $N 4
time GOMAXPROCS=4 go_to_c_concurrent/main $N 4

time GOMAXPROCS=8 go_concurrent/main $N 1
time GOMAXPROCS=8 go_to_c_concurrent/main $N 1

time GOMAXPROCS=8 go_concurrent/main $N 80000
time GOMAXPROCS=8 go_to_c_concurrent/main $N 80000

time GOMAXPROCS=8 go_concurrent/main $N 800000

# GOMAXPROCS=8 go_to_c_concurrent/main 8000000000 800000 &> out.txt
# Start of out.txt :
#
# Go wth 800000 goroutines calling C and GOMAXPROCS=8:
# Sum for n=1..8000000000 of 1/n =
# runtime/cgo: pthread_create failed: Resource temporarily unavailable
# SIGABRT: abort
# PC=0x7fff8791d0ae m=2

# goroutine 0 [idle]:

# goroutine 1 [runnable]:
# sync.(*WaitGroup).Wait(0x4192fd0)
#   /usr/local/go/src/sync/waitgroup.go:99
# main.main()
#   /Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-interlang/benchmarks/go_to_c_concurrent/main.go:39 +0x41c

# goroutine 17 [syscall, locked to thread]:
# runtime.goexit()
#   /usr/local/go/src/runtime/asm_amd64.s:1696 +0x1

# goroutine 138214 [runnable, locked to thread]:
# main._Cfunc_harmonic_range(0x5252d441, 0x5252fb50, 0x3ede5e2cc45e113a)
#   github.com/draffensperger/go-interlang/benchmarks/go_to_c_concurrent/_obj/_cgo_gotypes.go:41 +0x3a
# main.calcPartialSum(0x21b84, 0x5252d441, 0x5252fb50)
#   /Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-interlang/benchmarks/go_to_c_concurrent/main.go:18 +0x58
# created by main.main
#   /Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-interlang/benchmarks/go_to_c_concurrent/main.go:37 +0x3f5
