all: main

main: foo.o bar.c
    gcc foo.o bar.c -o main

foo.o: foo.go
    gccgo -c foo.go -o foo.o -fgo-prefix=example

clean:
    rm -f main *.o
