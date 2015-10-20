.PHONY: clean

UNAME := $(shell uname)

ifeq ($(UNAME), Darwin) 
    DYLIB_EXT := .dylib
else
    DYLIB_EXT := .so
endif

go_adder/libadder$(DYLIB_EXT):
	$(MAKE) -C go_adder

clean:
	$(MAKE) -C go_adder clean
