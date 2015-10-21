.PHONY: clean all

all: c_static_lib go_executable

c_static_lib:
	gcc -c greetings/*.c
	ar rs greetings.a *.o
	rm -rf *.o

go_executable:
	go build -o static_c_lib

clean:
	rm -rf *.o *.a static_c_lib
