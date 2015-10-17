cd adder
go build -buildmode=c-archive -o libadder.a
gcc -o main main.c adder/libadder.a


stategy:
- trim down the http example until it doesn't work anymore
- OK

go build -buildmode=c-shared -o libadder.a

gcc main.c -L ./adder -ladder

gcc -o main main.c adder/libadder.a

gcc -o main main.c adder/libadder.a -framework CoreFoundation -framework Security -lpthread


http://stackoverflow.com/questions/30896892/using-go-1-5-buildmode-c-archive-with-net-http-server-linked-from-c

go install -buildmode=c-shared

go build -buildmode=c-archive -o libadder.a adder/adder.go


go build -buildmode=c-shared -o libadder.a

go build -buildmode=c-shared -o adder.so

gcc main.c -L./adder -ladder

gcc main.c adder/adder.a

gcc main.c $GOPATH/pkg/darwin_amd64/github.com/draffensperger/go-inter-lang-calls/go_from_c/adder.a -framework CoreFoundation -framework Security -lpthread


stat $GOPATH/pkg/darwin_amd64/github.com/draffensperger/go-inter-lang-calls/go_from_c/adder.a


cd $GOPATH/pkg/darwin_amd64/github.com/draffensperger/go-inter-lang-calls/go_from_c

options:
- have Go have the main function
- use gccgo
- compile as a shared library

http://stackoverflow.com/questions/6125683/call-go-functions-from-c

brew install gcc

brew tap homebrew/versions

http://cxwangyi.github.io/notes/2014-04-26-build-gccgo-49-mac-os-x.html

http://cheesesun.blogspot.com/2010/04/callbacks-in-cgo.html


plan:
- start with a go main function
- compile with go and include c and calls back to go


https://github.com/shazow/gohttplib

https://groups.google.com/forum/#!topic/golang-nuts/UGYTlQFCGws

https://github.com/fiorix/gocp
