UNAME := $(shell uname)

ifeq ($(UNAME), Darwin) 
    DYLIB_EXT := .dylib
    FLAGS := -ldflags -s
else
    DYLIB_EXT := .so
    FLAGS := ''
endif

libadder$(DYLIB_EXT):
	# New in Go 1.5, build Go dynamic lib
	go build $(FLAGS) -o $@ -buildmode=c-shared

clean:
	rm -f libadder.*
