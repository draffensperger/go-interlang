
go build -buildmode=c-shared -o libadder.a

go install -buildmode=c-shared

go build -buildmode=c-archive -o libadder.a adder/adder.go


go build -buildmode=c-shared -o libadder.a

go build -buildmode=c-shared -o adder.so

gcc main.c -L./adder -ladder

gcc main.c adder/adder.a

gcc main.c $GOPATH/pkg/darwin_amd64/github.com/draffensperger/go-inter-lang-calls/go_from_c/adder.a -framework CoreFoundation -framework Security -lpthread


stat $GOPATH/pkg/darwin_amd64/github.com/draffensperger/go-inter-lang-calls/go_from_c/adder.a


cd $GOPATH/pkg/darwin_amd64/github.com/draffensperger/go-inter-lang-calls/go_from_c

