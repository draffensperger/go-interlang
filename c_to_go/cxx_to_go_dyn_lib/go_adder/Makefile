UNAME := $(shell uname)

ifeq ($(UNAME), Darwin) 
    DYLIB_EXT := .dylib
else
    DYLIB_EXT := .so
endif

libadder$(DYLIB_EXT):
	# New in Go 1.5, build Go dynamic lib
	go build -o $@ -buildmode=c-shared

clean:
	rm -f libadder.*
