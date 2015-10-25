# Benchmark results

Results run on a 2015 MacBook Pro with Intel with 2.8 GHz Intel Core i7 (Quad
Core with hyperthreading, Turbo Boost up to 4.0 Ghz). Running OS X 10.11 El 
Capitan. C compiler is Apple LLVM version 7.0.0 (clang-700.0.72). Used Go 1.5
and Ruby 2.2.3.

## Harmonic series benchmarks

### Go to C call benchmarks

Output of `./go_to_c.sh`:

```
=======================================
--- Go to C interlang call cost ---
=======================================
size of long: 8

C only:
Sum for n=1..800000000 of 1/n =
21.077337951159535
real    0m2.921s
user    0m2.916s
sys     0m0.003s


Go only:
Sum for n=1..800000000 of 1/n =
21.077337951159535
real    0m5.422s
user    0m5.415s
sys     0m0.008s


Go with 1 time C harmonic_sum call:
Sum for n=1..800000000 of 1/n =
21.077337951159535
real    0m2.955s
user    0m2.948s
sys     0m0.004s


Go with n times C add_harmonic calls:
Sum for n=1..800000000 of 1/n =
21.077337951159535
real    1m56.478s
user    1m56.312s
sys     0m0.198s
```

### Go to Ruby call benchmarks

Output of `./ruby_to_go.sh`:

```
=======================================
--- Ruby to Go FFI cost ---
=======================================


C only:
Sum for n=1..80000000 of 1/n =
18.774752863787992
real    0m0.304s
user    0m0.300s
sys     0m0.002s


Go only:
Sum for n=1..80000000 of 1/n =
18.77475286378799
real    0m0.553s
user    0m0.547s
sys     0m0.005s


Ruby only:
Sum for n=1..80000000 of 1/n =
18.77475286378799
real    0m5.453s
user    0m5.421s
sys     0m0.027s


Ruby with 1 time Go HarmonicSum call:
Sum for n=1..80000000 of 1/n =
18.77475286378799
real    0m0.662s
user    0m0.623s
sys     0m0.035s


Ruby with n times Go AddHarmonic calls:
Sum for n=1..80000000 of 1/n =
18.77475286378799
real    2m0.839s
user    1m12.624s
sys     0m48.176s


Ruby to Go no-op FFI call benchmark:
       user     system      total        real
  48.200000  30.660000  78.860000 ( 78.977828)
Ruby function call benchmark:
       user     system      total        real
   2.830000   0.010000   2.840000 (  2.839761)
```

### Concurrency benchmarks

Output of `./concurrency.sh`:

```
=======================================
--- Go to C and concurrency ---
=======================================


Go only:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m56.482s
user    0m56.253s
sys     0m0.122s


C only:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m29.354s
user    0m29.284s
sys     0m0.037s


Go with 1 goroutines and GOMAXPROCS=1:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m55.320s
user    0m55.204s
sys     0m0.103s


Go wth 1 goroutines calling C and GOMAXPROCS=1:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m29.568s
user    0m29.463s
sys     0m0.055s


Go with 8 goroutines and GOMAXPROCS=1:
Sum for n=1..8000000000 of 1/n =
23.379923043590928
real    0m55.868s
user    0m55.725s
sys     0m0.134s


Go wth 8 goroutines calling C and GOMAXPROCS=1:
Sum for n=1..8000000000 of 1/n =
23.379923043590928
real    0m7.479s
user    0m57.470s
sys     0m0.092s


Go with 8 goroutines and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.379923043590928
real    0m7.858s
user    1m0.744s
sys     0m0.091s


Go wth 8 goroutines calling C and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.379923043590928
real    0m7.610s
user    0m58.936s
sys     0m0.098s


Go with 4 goroutines and GOMAXPROCS=4:
Sum for n=1..8000000000 of 1/n =
23.379923043595255
real    0m14.466s
user    0m57.521s
sys     0m0.116s


Go wth 4 goroutines calling C and GOMAXPROCS=4:
Sum for n=1..8000000000 of 1/n =
23.379923043595255
real    0m8.057s
user    0m31.979s
sys     0m0.066s


Go with 1 goroutines and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m56.975s
user    0m56.851s
sys     0m0.115s


Go wth 1 goroutines calling C and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m29.049s
user    0m29.005s
sys     0m0.028s


Go with 80000 goroutines and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.37992304358969
real    0m7.633s
user    0m59.654s
sys     0m0.107s


Go wth 80000 goroutines calling C and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.37992304358969
real    0m7.492s
user    0m58.981s
sys     0m0.199s


Go with 800000 goroutines and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.37992304358886
real    0m9.613s
user    1m5.301s
sys     0m0.897s

Go wth 800000 goroutines calling C and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
runtime/cgo: pthread_create failed: Resource temporarily unavailable
SIGABRT: abort
PC=0x7fff8791d0ae m=2

goroutine 0 [idle]:

goroutine 1 [runnable]:
sync.(*WaitGroup).Wait(0x4192fd0)
  /usr/local/go/src/sync/waitgroup.go:99
main.main()
  /Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-interlang/benchmarks/go_to_c_concurrent/main.go:39 +0x41c

goroutine 17 [syscall, locked to thread]:
runtime.goexit()
  /usr/local/go/src/runtime/asm_amd64.s:1696 +0x1

goroutine 138214 [runnable, locked to thread]:
main._Cfunc_harmonic_range(0x5252d441, 0x5252fb50, 0x3ede5e2cc45e113a)
  github.com/draffensperger/go-interlang/benchmarks/go_to_c_concurrent/_obj/_cgo_gotypes.go:41 +0x3a
main.calcPartialSum(0x21b84, 0x5252d441, 0x5252fb50)
  /Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-interlang/benchmarks/go_to_c_concurrent/main.go:18 +0x58
created by main.main
  /Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-interlang/benchmarks/go_to_c_concurrent/main.go:37 +0x3f5

... many more goroutines printing status, aborted output ...
```

### HTTP call benchmarks

```
=======================================
--- Ruby to Go localhost HTTP cost ---
=======================================


C only:
Sum for n=1..16000 of 1/n =
10.257590915797937
real    0m0.003s
user    0m0.001s
sys     0m0.001s


Go only:
Sum for n=1..16000 of 1/n =
10.257590915797937
real    0m0.005s
user    0m0.001s
sys     0m0.003s


Ruby only:
Sum for n=1..16000 of 1/n =
10.257590915797937
real    0m0.070s
user    0m0.042s
sys     0m0.025s


Ruby with n times Go AddHarmonic calls:
Sum for n=1..16000 of 1/n =
10.257590915797937
real    0m0.155s
user    0m0.106s
sys     0m0.045s


Go with n times C add_harmonic calls:
Sum for n=1..16000 of 1/n =
10.257590915797937
real    0m0.007s
user    0m0.003s
sys     0m0.003s


Go with n times HTTP calls:
Sum for n=1..16000 of 1/n =
10.257590915797937
real    0m1.181s
user    0m0.522s
sys     0m0.398s


Ruby with n times HTTP calls:
Sum for n=1..16000 of 1/n =
10.257590915797937
real    0m9.424s
user    0m2.586s
sys     0m1.660s
```
