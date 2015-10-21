.PHONY: clean

CCFLAGS := -pthread
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S), Darwin)
	# This prevents a warning about PIE on Mac OS
	CCFLAGS += -Wl,-no_pie
endif

static_go_lib: adder/libadder.a
	gcc $(CCFLAGS) -o $@ *.c $<

adder/libadder.a:
	$(MAKE) -C adder

clean:
	rm -f adder/libadder.a static_go_lib
