
cd adder
go build -buildmode=c-archive -o libadder.a
gcc -o main main.c adder/libadder.a

gcc main.c -L ./adder -ladder
gcc -o main main.c adder/libadder.a
gcc -o main main.c adder/libadder.a -framework CoreFoundation -framework Security -lpthread
