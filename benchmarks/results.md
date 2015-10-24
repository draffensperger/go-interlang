# Benchmark results

Results run on a 2015 MacBook Pro with Intel with 2.8 GHz Intel Core i7 (Quad
Core with hyperthreading, Turbo Boost up to 4.0 Ghz). Running OS X 10.11 El 
Capitan. C compiler is Apple LLVM version 7.0.0 (clang-700.0.72). Go 1.5. 
Ruby 2.2.3.

## Go to C call benchmarks

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

## Go to Ruby call benchmarks

Output of `./ruby_to_go.sh`:

```
=======================================
--- Ruby to Go FFI cost ---
=======================================


C only:
Sum for n=1..80000000 of 1/n =
18.774752863787992
real    0m0.300s
user    0m0.296s
sys     0m0.001s


Go only:
Sum for n=1..80000000 of 1/n =
18.77475286378799
real    0m0.571s
user    0m0.562s
sys     0m0.006s


Ruby only:
Sum for n=1..80000000 of 1/n =
18.77475286378799
real    0m3.839s
user    0m3.801s
sys     0m0.032s


Ruby with 1 time Go HarmonicSum call:
Sum for n=1..80000000 of 1/n =
18.77475286378799
real    0m0.674s
user    0m0.633s
sys     0m0.036s


Ruby with n times Go AddHarmonic calls:
Sum for n=1..80000000 of 1/n =
18.77475286378799
real    2m7.797s
user    1m17.421s
sys     0m50.121s
```

## Concurrency benchmarks

Output of `./concurrency.sh`:

```
=======================================
--- Go to C and concurrency ---
=======================================


Go only:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m55.671s
user    0m55.576s
sys     0m0.094s
size of long: 8

C only:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m29.448s
user    0m29.408s
sys     0m0.024s


Go with 1 goroutines and GOMAXPROCS=1:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m54.353s
user    0m54.318s
sys     0m0.057s


Go wth 1 goroutines calling C and GOMAXPROCS=1:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m29.437s
user    0m29.384s
sys     0m0.031s


Go with 8 goroutines and GOMAXPROCS=1:
Sum for n=1..8000000000 of 1/n =
23.379923043590928
real    0m55.073s
user    0m54.942s
sys     0m0.112s


Go wth 8 goroutines calling C and GOMAXPROCS=1:
Sum for n=1..8000000000 of 1/n =
23.379923043590928
real    0m7.937s
user    1m1.193s
sys     0m0.106s


Go with 8 goroutines and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.379923043590928
real    0m8.681s
user    1m7.312s
sys     0m0.066s


Go wth 8 goroutines calling C and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.379923043590928
real    0m8.345s
user    1m3.959s
sys     0m0.093s


Go with 4 goroutines and GOMAXPROCS=4:
Sum for n=1..8000000000 of 1/n =
23.379923043595255
real    0m15.340s
user    1m1.270s
sys     0m0.040s


Go wth 4 goroutines calling C and GOMAXPROCS=4:
Sum for n=1..8000000000 of 1/n =
23.379923043595255
real    0m8.674s
user    0m34.459s
sys     0m0.027s


Go with 1 goroutines and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m56.440s
user    0m56.278s
sys     0m0.128s


Go wth 1 goroutines calling C and GOMAXPROCS=8:
Sum for n=1..8000000000 of 1/n =
23.379923043614358
real    0m31.438s
user    0m31.270s
sys     0m0.082s
```

## HTTP call benchmarks

```
=======================================
--- Ruby to Go localhost HTTP cost ---
=======================================


C only:
Sum for n=1..10000 of 1/n =
9.787606036044348
real    0m0.003s
user    0m0.001s
sys     0m0.001s


Go only:
Sum for n=1..10000 of 1/n =
9.787606036044348
real    0m0.005s
user    0m0.001s
sys     0m0.003s


Ruby only:
Sum for n=1..10000 of 1/n =
9.787606036044348
real    0m0.073s
user    0m0.042s
sys     0m0.027s


Ruby with n times Go AddHarmonic calls:
Sum for n=1..10000 of 1/n =
9.787606036044348
real    0m0.146s
user    0m0.099s
sys     0m0.041s


Go with n times HTTP calls:
Sum for n=1..10000 of 1/n =
9.787606036044348
real    0m0.746s
user    0m0.327s
sys     0m0.259s


Ruby with n times HTTP calls:
Sum for n=1..10000 of 1/n =
9.787606036044348
real    0m5.785s
user    0m1.635s
sys     0m1.042s
```
