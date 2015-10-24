examples still to do:

ruby to go 1 http
ruby to go 1 http's

# wonder if this shows that c routines can count as syscalls
# shows that go is not stuck in a sleepy c routine

http://stackoverflow.com/questions/14818084/what-is-the-proper-include-for-the-function-sleep-in-c

go to c calls max cost:
(60+56)/800000000
= 1.45e-7
600 cycles


1/4e9 is time for one cycle

145 nanoseconds at the most


millisecond e-3
microsecond e-6
nanosecond e-9

14.5 e-8
K e3
M e6
G e9
4.0 Ghz



N=4000000000

N=400000000
/usr/bin/time -l single_lang/go_only $N
/usr/bin/time -l single_lang/c_only $N
/usr/bin/time -l ruby single_lang/ruby_only.rb $N

single langs:
c 0
ruby 9.79
go 0.11

time go_to_c/go_to_c $N

# 

go to c adder

ruby only
ruby to go adder ffi
ruby to go adder http

# concurrency
go pausing number server
c socket server fetcher w/ go routines
go socket server fetcher w/ go routiens


https://github.com/nbraud/benchmarksgame

http://benchmarksgame.alioth.debian.org/

http://pybenchmarks.org/u64q/performance.php?test=spectralnorm

# how to limit concurrency of the go fetcher
http://jmoiron.net/blog/limiting-concurrency-in-go/

#
http://stackoverflow.com/questions/22077802/simple-c-example-of-doing-an-http-post-and-consuming-the-response

http://stackoverflow.com/questions/28354141/c-code-and-goroutine-scheduling

https://github.com/golang/go/issues/8636
https://groups.google.com/forum/#!topic/golang-nuts/9RqGvk-b3Ow

https://groups.google.com/forum/#!topic/golang-nuts/RTtMsgZi88Q

http://grokbase.com/t/gg/golang-nuts/12a9ezsvg3/go-nuts-c-goroutines-abysmal-performance

http://osdir.com/ml/go-language-discuss/2013-08/msg03029.html

http://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/

https://groups.google.com/forum/#!topic/golang-nuts/RTtMsgZi88Q
