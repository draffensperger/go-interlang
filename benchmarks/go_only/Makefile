UNAME := $(shell uname)
ifeq ($(UNAME), Darwin) 
    DYLIB_EXT := .dylib
else
    DYLIB_EXT := .so
endif

all:
	go build -o mai
	go build -o libharmonic$(DYLIB_EXT) -buildmode=c-shared
