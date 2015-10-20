.PHONY: clean

c_to_static_lib: adder/libadder.a
	gcc -Wl,-no_pie -o $@ *.c $<

adder/libadder.a:
	$(MAKE) -C adder

clean:
	rm -f adder/libadder.a c_to_static_go_lib
