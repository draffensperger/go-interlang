UNAME := $(shell uname)

ifeq ($(UNAME), Darwin) 
    DYLIB_EXT := .dylib
else
    DYLIB_EXT := .so
endif

libadder$(DYLIB_EXT):
	go build -o $@ -buildmode=c-shared

clean:
	rm -f libadder.*
