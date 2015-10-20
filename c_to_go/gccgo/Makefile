all: main

main: main.o adder.o
	gccgo *.o -o $@

main.o: main.c
	gcc -c $< -o $@

adder.o: adder.go
	gccgo -c $< -o $@ -fgo-prefix=example

clean:
	rm -f main *.o
