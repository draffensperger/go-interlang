.PHONY: clean

UNAME := $(shell uname)

ifeq ($(UNAME), Darwin) 
    DYLIB_EXT := .dylib
else
    DYLIB_EXT := .so
endif

all: go_adder/libadder$(DYLIB_EXT) main 

go_adder/libadder$(DYLIB_EXT):
	$(MAKE) -C go_adder

main:
	g++ -o $@ main.cxx -I./go_adder -L./go_adder -ladder

clean:
	rm -f main go_adder/libadder.*
