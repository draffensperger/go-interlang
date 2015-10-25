UNAME := $(shell uname)
ifeq ($(UNAME), Darwin) 
    DYLIB_EXT := .dylib
else
    DYLIB_EXT := .so
endif

libgo_no_op:
	go build -o libgo_no_op$(DYLIB_EXT) -buildmode=c-shared
